package controller

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const targetDurationSec = 30

type AWCJob struct {
	controller       *AWC
	awc              *models.AutoWaterChange
	freshPump        *models.Pump
	wastePump        *models.Pump
	freshFirmata     *gomata.Firmata
	wasteFirmata     *gomata.Firmata
	freshCalibration *models.Calibration
	wasteCalibration *models.Calibration
}

func NewAWCJob(
	controller *AWC,
	awc *models.AutoWaterChange,
	freshPump, wastePump *models.Pump,
	freshFirmata, wasteFirmata *gomata.Firmata,
	freshCalibration, wasteCalibration *models.Calibration,
) *AWCJob {
	return &AWCJob{
		controller:       controller,
		awc:              awc,
		freshPump:        freshPump,
		wastePump:        wastePump,
		freshFirmata:     freshFirmata,
		wasteFirmata:     wasteFirmata,
		freshCalibration: freshCalibration,
		wasteCalibration: wasteCalibration,
	}
}

func (j *AWCJob) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Starting AWC job %s", j.awc.ID)

	var (
		// Exchange rate is in L/day, stepper is in steps/sec
		freshMlPerSecond = (j.awc.ExchangeRate*1000 + j.awc.SalinityAdjustment) / 24 / 60 / 60
		wasteMlPerSecond = (j.awc.ExchangeRate*1000 - j.awc.SalinityAdjustment) / 24 / 60 / 60
	)

	// Handle salinity adjustment greater than the change amount gracefully
	if freshMlPerSecond < 0 {
		freshMlPerSecond = 0
	}
	if wasteMlPerSecond < 0 {
		wasteMlPerSecond = 0
	}

	err := j.freshFirmata.StepperZero(int(j.freshPump.DeviceID))
	if err != nil {
		j.reset(err)
		j.event(AWCJobErrorKind, "Failure zeroing fresh pump: %s", err)
	}

	// NOTE: Target duration is (essentially) the duration of time the pump will
	// be commanded to pump for. To ensure a constant pump speed, the ticker could
	// fire faster, causing the next position to be scheduled before the previous
	// dose completed.
	ticker := time.NewTicker(targetDurationSec * time.Second)
	defer ticker.Stop()

	// Run first job immediately
	jobCtx, jobCancel := context.WithTimeout(ctx, 2*targetDurationSec*time.Second)
	j.runJob(jobCtx, freshMlPerSecond, wasteMlPerSecond)
	jobCancel()

	for {
		select {
		case <-ticker.C:
			jobCtx, jobCancel := context.WithTimeout(ctx, 2*targetDurationSec*time.Second)
			j.runJob(jobCtx, freshMlPerSecond, wasteMlPerSecond)
			jobCancel()

		case <-ctx.Done():
			return
		}
	}
}

func (j *AWCJob) event(kind string, data string, args ...interface{}) {
	j.controller.eventCh <- &models.AwcEvent{
		ID:                uuid.New().String(),
		AutoWaterChangeID: j.awc.ID,
		Timestamp:         time.Now().Unix(),
		Kind:              kind,
		Data:              fmt.Sprintf(data, args...),
	}
}

func (j *AWCJob) runJob(ctx context.Context, freshMlPerSecond, wasteMlPerSecond float64) {
	freshDone, status, err := j.dose(ctx, "fresh", j.freshFirmata, j.freshPump, j.freshCalibration, freshMlPerSecond)
	if err == context.DeadlineExceeded || errors.Is(err, io.EOF) {
		j.reset(err)
		return
	}
	if err != nil {
		j.event(status, err.Error())
		if errors.Is(err, gomata.ErrNotConnected) {
			j.controller.Reset()
		}
		return
	}

	wasteDone, status, err := j.dose(ctx, "waste", j.wasteFirmata, j.wastePump, j.wasteCalibration, wasteMlPerSecond)
	if err == context.DeadlineExceeded || errors.Is(err, io.EOF) {
		j.reset(err)
		return
	}
	if err != nil {
		j.event(status, err.Error())
		return
	}

	select {
	case report := <-freshDone:
		volume := float64(report.Position) * j.freshCalibration.Volume / float64(j.freshCalibration.Steps)
		j.recordDose(ctx, j.freshPump, volume, "AWC fresh pump")
	case <-ctx.Done():
		// If we timed out, reconnect to firmata and recreate the ATO jobs
		if ctx.Err() == context.DeadlineExceeded {
			j.reset(ctx.Err())
			return
		}

		log.Printf("Job context cancelled, terminating")
		return
	}

	select {
	case report := <-wasteDone:
		volume := float64(report.Position) * j.wasteCalibration.Volume / float64(j.wasteCalibration.Steps)
		j.recordDose(ctx, j.wastePump, volume, "AWC waste pump")
	case <-ctx.Done():
		// If we timed out, reconnect to firmata and recreate the ATO jobs
		if ctx.Err() == context.DeadlineExceeded {
			j.reset(ctx.Err())
			return
		}

		log.Printf("Job context cancelled, terminating")
		return
	}
}

func (j *AWCJob) dose(ctx context.Context, name string, firmata *gomata.Firmata, pump *models.Pump, calibration *models.Calibration, mlPerSecond float64) (<-chan gomata.StepperPosition, string, error) {
	err := firmata.StepperZero(int(pump.DeviceID))
	if err != nil {
		return nil, AWCJobErrorKind, fmt.Errorf("Failure zeroing %s pump speed (aborting job run): %w", name, err)
	}

	var (
		now         = time.Now()
		nextTime    = now.Add(targetDurationSec * time.Second)
		nextElapsed = nextTime.Sub(now)
		nextSteps   = mlPerSecond * float64(nextElapsed/time.Second) * float64(calibration.Steps) / calibration.Volume
		speed       = roundedSpeed(nextSteps / float64(targetDurationSec))
	)

	if pump.Acceleration.Valid {
		err = firmata.StepperSetAcceleration(int(pump.DeviceID), float32(pump.Acceleration.Float64))
		if err != nil {
			return nil, AWCJobErrorKind, fmt.Errorf("Failure setting %s pump speed (aborting job run): %w", name, err)
		}
	}

	err = firmata.StepperSetSpeed(int(pump.DeviceID), float32(speed))
	if err != nil {
		return nil, AWCJobErrorKind, fmt.Errorf("Failure setting %s pump speed (aborting job run): %w", name, err)
	}

	reportCh := firmata.AwaitStepperMoveCompletion(int32(pump.DeviceID))

	err = firmata.StepperStep(int(pump.DeviceID), int32(nextSteps))
	if err != nil {
		return nil, AWCJobErrorKind, fmt.Errorf("stepping %s pump (aborting job run): %w", name, err)
	}

	return reportCh, "", nil
}

func (j *AWCJob) getPosition(ctx context.Context, firmata *gomata.Firmata, pump *models.Pump) (*gomata.StepperPosition, error) {
	reportCh := firmata.AwaitStepperReport(int32(pump.DeviceID))

	err := firmata.StepperReport(int(pump.DeviceID))
	if err != nil {
		return nil, fmt.Errorf("requesting pump position report: %w", err)
	}

	select {
	case pos := <-reportCh:
		return &pos, nil

	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timed out waiting for pump position report")

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (j *AWCJob) recordDose(ctx context.Context, pump *models.Pump, volume float64, message string, args ...interface{}) {
	dose := models.Dose{
		ID:        uuid.New().String(),
		PumpID:    pump.ID,
		Timestamp: time.Now().Unix(),
		Volume:    volume,
		Message:   null.StringFrom(fmt.Sprintf(message, args...)),
	}
	err := dose.Insert(ctx, j.controller.db, boil.Infer())
	if err != nil {
		j.event(AWCJobErrorKind, "Failure to insert dose: %s", err)
	}
	log.Printf("%s pumped %fmL", pump.Name.String, volume)
}

func (j *AWCJob) reset(err error) {
	log.Printf("Resetting AWC job: %s", err)

	panic("AWC reset")
}
