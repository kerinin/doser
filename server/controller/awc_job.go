package controller

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

const targetDurationSec = 60

type AWCJob struct {
	eventCh          chan<- Event
	awc              *models.AutoWaterChange
	freshPump        *models.Pump
	wastePump        *models.Pump
	freshFirmata     *gomata.Firmata
	wasteFirmata     *gomata.Firmata
	freshCalibration *models.Calibration
	wasteCalibration *models.Calibration
}

func NewAWCJob(
	eventCh chan<- Event,
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

	for {
		select {
		case <-ticker.C:
			currentFreshPosition, err := j.getPosition(ctx, j.freshFirmata, j.freshPump)
			if err != nil {
				j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("getting fresh pump position (aborting job run): %w", err)}
				continue
			}
			currentWastePosition, err := j.getPosition(ctx, j.wasteFirmata, j.wastePump)
			if err != nil {
				j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("getting fresh pump position (aborting job run): %w", err)}
				continue
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
				j.eventCh <- &AWCStatus{j.awc, time.Now().Sub(initialTime), totalFreshVolume, totalWasteVolume}
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

			err = j.freshFirmata.StepperSetSpeed(int(j.freshPump.DeviceID), float32(freshSpeed))
			if err != nil {
				j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("setting fresh pump speed (aborting job run): %w", err)}
				continue
			}
			err = j.wasteFirmata.StepperSetSpeed(int(j.wastePump.DeviceID), float32(wasteSpeed))
			if err != nil {
				j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("setting waste pump speed (aborting job run): %w", err)}
				continue
			}

			err = j.freshFirmata.StepperTo(int(j.freshPump.DeviceID), int32(nextFreshSteps))
			if err != nil {
				j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("stepping fresh pump (aborting job run): %w", err)}
				continue
			}
			err = j.wasteFirmata.StepperTo(int(j.wastePump.DeviceID), int32(nextWasteSteps))
			if err != nil {
				j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("stepping waste pump (aborting job run): %w", err)}
				continue
			}

		case <-ctx.Done():
			err := j.freshFirmata.StepperStop(int(j.freshPump.DeviceID))
			if err != nil {
				j.eventCh <- &UncontrolledPumpError{j.freshPump.ID, fmt.Errorf("stopping fresh pump during shutdown of ATO job: %w", err)}
			}
			err = j.wasteFirmata.StepperStop(int(j.wastePump.DeviceID))
			if err != nil {
				j.eventCh <- &UncontrolledPumpError{j.wastePump.ID, fmt.Errorf("stopping waste pump during shutdown of ATO job: %w", err)}
			}

			return
		}
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
