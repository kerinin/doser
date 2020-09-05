package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const targetDurationSec = 60

type AWCJob struct {
	eventCh          chan<- *models.AwcEvent
	db               *sql.DB
	awc              *models.AutoWaterChange
	freshPump        *models.Pump
	wastePump        *models.Pump
	freshFirmata     *gomata.Firmata
	wasteFirmata     *gomata.Firmata
	freshCalibration *models.Calibration
	wasteCalibration *models.Calibration
}

func NewAWCJob(
	eventCh chan<- *models.AwcEvent,
	db *sql.DB,
	awc *models.AutoWaterChange,
	freshPump, wastePump *models.Pump,
	freshFirmata, wasteFirmata *gomata.Firmata,
	freshCalibration, wasteCalibration *models.Calibration,
) *AWCJob {
	return &AWCJob{
		eventCh:          eventCh,
		db:               db,
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
		mlPerSecond = j.awc.ExchangeRate * 1000 / 24 / 60 / 60
	)

	err := j.freshFirmata.StepperZero(int(j.freshPump.DeviceID))
	if err != nil {
		j.event(AWCJobErrorKind, "Failure zeroing fresh pump: %w", err)
	}

	ticker := time.NewTicker(targetDurationSec * time.Second)
	defer ticker.Stop()

	// Run first job immediately
	j.runJob(ctx, mlPerSecond)

	for {
		select {
		case <-ticker.C:
			status, err := j.runJob(ctx, mlPerSecond)
			if err != nil {
				j.event(status, err.Error())
			}

		case <-ctx.Done():
			err := j.freshFirmata.StepperStop(int(j.freshPump.DeviceID))
			if err != nil {
				j.event(UncontrolledPumpKind, "Failure stopping fresh pump %s during shutdown of AWC job: %w", j.freshPump.ID, err)
			}
			err = j.wasteFirmata.StepperStop(int(j.wastePump.DeviceID))
			if err != nil {
				j.event(UncontrolledPumpKind, "Failure stopping waste pump %s during shutdown of AWC job: %w", j.wastePump.ID, err)
			}

			return
		}
	}
}

func (j *AWCJob) event(kind string, data string, args ...interface{}) {
	j.eventCh <- &models.AwcEvent{
		ID:                uuid.New().String(),
		AutoWaterChangeID: j.awc.ID,
		Timestamp:         time.Now().Unix(),
		Kind:              kind,
		Data:              fmt.Sprintf(data, args...),
	}
}

func (j *AWCJob) runJob(ctx context.Context, mlPerSecond float64) (string, error) {
	status, err := j.dose(ctx, "fresh", j.freshFirmata, j.freshPump, j.freshCalibration, mlPerSecond)
	if err != nil {
		return status, err
	}

	status, err = j.dose(ctx, "waste", j.wasteFirmata, j.wastePump, j.wasteCalibration, mlPerSecond)
	if err != nil {
		return status, err
	}

	return "", nil
}

func (j *AWCJob) dose(ctx context.Context, name string, firmata *gomata.Firmata, pump *models.Pump, calibration *models.Calibration, mlPerSecond float64) (string, error) {
	report, err := j.getPosition(ctx, firmata, pump)
	if err != nil {
		return AWCJobErrorKind, fmt.Errorf("Failure getting %s pump position (aborting job run): %w", name, err)
	}
	if report.Position > 0 {
		volume := float64(report.Position) * calibration.Volume / float64(calibration.Steps)
		j.recordDose(ctx, pump, volume, "AWC %s pump", name)
	}

	err = firmata.StepperZero(int(pump.DeviceID))
	if err != nil {
		return AWCJobErrorKind, fmt.Errorf("Failure zeroing %s pump speed (aborting job run): %w", name, err)
	}

	var (
		nextTime    = time.Now().Add(targetDurationSec * time.Second)
		nextElapsed = nextTime.Sub(time.Now())
		nextSteps   = mlPerSecond * float64(nextElapsed/time.Second) * float64(calibration.Steps) / calibration.Volume
		speed       = nextSteps / targetDurationSec
	)

	if pump.Acceleration.Valid {
		err = firmata.StepperSetAcceleration(int(pump.DeviceID), float32(pump.Acceleration.Float64))
		if err != nil {
			return AWCJobErrorKind, fmt.Errorf("Failure setting %s pump speed (aborting job run): %w", name, err)
		}
	}

	err = firmata.StepperSetSpeed(int(pump.DeviceID), float32(math.Floor(speed)))
	if err != nil {
		return AWCJobErrorKind, fmt.Errorf("Failure setting %s pump speed (aborting job run): %w", name, err)
	}

	err = firmata.StepperTo(int(pump.DeviceID), int32(nextSteps))
	if err != nil {
		return AWCJobErrorKind, fmt.Errorf("stepping %s pump (aborting job run): %w", name, err)
	}

	return "", nil
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
	err := dose.Insert(ctx, j.db, boil.Infer())
	if err != nil {
		j.event(ATOJobErrorKind, "Failure to insert dose: %w", err)
	}
}
