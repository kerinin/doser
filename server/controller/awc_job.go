package controller

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

const targetDurationSec = 60

type AWCJob struct {
	eventCh          chan<- model.AutoWaterChangeEvent
	awc              *models.AutoWaterChange
	freshPump        *models.Pump
	wastePump        *models.Pump
	freshFirmata     *gomata.Firmata
	wasteFirmata     *gomata.Firmata
	freshCalibration *models.Calibration
	wasteCalibration *models.Calibration
}

func NewAWCJob(
	eventCh chan<- model.AutoWaterChangeEvent,
	awc *models.AutoWaterChange,
	freshPump, wastePump *models.Pump,
	freshFirmata, wasteFirmata *gomata.Firmata,
	freshCalibration, wasteCalibration *models.Calibration,
) *AWCJob {
	return &AWCJob{
		eventCh:          eventCh,
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
		initialTime = time.Now()

		// Exchange rate is in L/day, stepper is in steps/sec
		mlPerSecond          = j.awc.ExchangeRate * 1000 / 24 / 60 / 60
		initialFreshPosition *gomata.StepperPosition
		initialWastePosition *gomata.StepperPosition
	)

	ticker := time.NewTicker(targetDurationSec * time.Second)
	defer ticker.Stop()

	// Run first job immediately
	j.runJob(ctx, initialTime, mlPerSecond, initialFreshPosition, initialWastePosition)

	for {
		select {
		case <-ticker.C:
			j.runJob(ctx, initialTime, mlPerSecond, initialFreshPosition, initialWastePosition)

		case <-ctx.Done():
			err := j.freshFirmata.StepperStop(int(j.freshPump.DeviceID))
			if err != nil {
				j.eventCh <- &model.UncontrolledPumpError{int(time.Now().Unix()), j.freshPump.ID, fmt.Errorf("stopping fresh pump during shutdown of ATO job: %w", err).Error()}
			}
			err = j.wasteFirmata.StepperStop(int(j.wastePump.DeviceID))
			if err != nil {
				j.eventCh <- &model.UncontrolledPumpError{int(time.Now().Unix()), j.wastePump.ID, fmt.Errorf("stopping waste pump during shutdown of ATO job: %w", err).Error()}
			}

			return
		}
	}
}

func (j *AWCJob) runJob(ctx context.Context, initialTime time.Time, mlPerSecond float64, initialFreshPosition, initialWastePosition *gomata.StepperPosition) {
	currentFreshPosition, err := j.getPosition(ctx, j.freshFirmata, j.freshPump)
	if err != nil {
		j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("getting fresh pump position (aborting job run): %w", err).Error()}
		return
	}
	currentWastePosition, err := j.getPosition(ctx, j.wasteFirmata, j.wastePump)
	if err != nil {
		j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("getting fresh pump position (aborting job run): %w", err).Error()}
		return
	}

	if initialFreshPosition == nil || initialWastePosition == nil {
		initialFreshPosition = currentFreshPosition
		initialWastePosition = currentWastePosition
	} else {
		var (
			totalFreshSteps  = currentFreshPosition.Position - initialFreshPosition.Position
			totalFreshVolume = float64(totalFreshSteps) * j.freshCalibration.Volume / float64(j.freshCalibration.Steps)
			totalWasteSteps  = currentWastePosition.Position - initialWastePosition.Position
			totalWasteVolume = float64(totalWasteSteps) * j.wasteCalibration.Volume / float64(j.wasteCalibration.Steps)
		)
		j.eventCh <- &model.AWCStatus{int(time.Now().Unix()), time.Now().Sub(initialTime).Seconds(), totalFreshVolume, totalWasteVolume}
	}

	// NOTE: This attempts to minimize accumulated inaccuracy due to
	// rounding errors or precision losses. At each step, it determines
	// when the current step will end, then uses the target exchange
	// rate to determine how much fluid should have been pumped by that
	// time. The speed is set to pump the required amount in the expected
	// duration and the pump is instructed to move to the target
	// volume (in steps).
	var (
		nextTime    = time.Now().Add(targetDurationSec * time.Second)
		nextElapsed = nextTime.Sub(initialTime)

		nextFreshSteps = mlPerSecond * float64(nextElapsed/time.Second) * float64(j.freshCalibration.Steps) / j.freshCalibration.Volume
		nextWasteSteps = mlPerSecond * float64(nextElapsed/time.Second) * float64(j.wasteCalibration.Steps) / j.wasteCalibration.Volume

		deltaFreshSteps = nextFreshSteps - float64(currentFreshPosition.Position)
		deltaWasteSteps = nextWasteSteps - float64(currentWastePosition.Position)

		freshSpeed = deltaFreshSteps / targetDurationSec
		wasteSpeed = deltaWasteSteps / targetDurationSec
	)

	// TODO: make sure we don't go backwards

	if j.freshPump.Acceleration.Valid {
		err = j.freshFirmata.StepperSetAcceleration(int(j.freshPump.DeviceID), float32(j.freshPump.Acceleration.Float64))
		if err != nil {
			j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("setting fresh pump speed (aborting job run): %w", err).Error()}
			return
		}
	}

	if j.wastePump.Acceleration.Valid {
		err = j.wasteFirmata.StepperSetAcceleration(int(j.wastePump.DeviceID), float32(j.wastePump.Acceleration.Float64))
		if err != nil {
			j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("setting waste pump speed (aborting job run): %w", err).Error()}
			return
		}
	}

	err = j.freshFirmata.StepperSetSpeed(int(j.freshPump.DeviceID), float32(math.Floor(freshSpeed)))
	if err != nil {
		j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("setting fresh pump speed (aborting job run): %w", err).Error()}
		return
	}
	err = j.wasteFirmata.StepperSetSpeed(int(j.wastePump.DeviceID), float32(math.Floor(wasteSpeed)))
	if err != nil {
		j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("setting waste pump speed (aborting job run): %w", err).Error()}
		return
	}

	err = j.freshFirmata.StepperTo(int(j.freshPump.DeviceID), int32(nextFreshSteps))
	if err != nil {
		j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("stepping fresh pump (aborting job run): %w", err).Error()}
		return
	}
	err = j.wasteFirmata.StepperTo(int(j.wastePump.DeviceID), int32(nextWasteSteps))
	if err != nil {
		j.eventCh <- &model.AWCJobError{int(time.Now().Unix()), fmt.Errorf("stepping waste pump (aborting job run): %w", err).Error()}
		return
	}
}

func (j *AWCJob) getPosition(ctx context.Context, firmata *gomata.Firmata, pump *models.Pump) (*gomata.StepperPosition, error) {
	reportCh := firmata.AwaitStepperReport(int32(pump.DeviceID))

	// TODO: The firmata hasn't been connected when we get here and it causes a panic
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
