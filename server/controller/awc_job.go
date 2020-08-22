package controller

import (
	"context"
	"fmt"
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
	awc *models.AutoWaterChange,
	freshPump, wastePump *models.Pump,
	freshFirmata, wasteFirmata *gomata.Firmata,
	freshCalibration, wasteCalibration *models.Calibration,
) *AWCJob {
	return &AWCJob{}
}

func (j *AWCJob) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		// Exchange rate is in L/day, stepper is in steps/sec
		mlPerSecond = j.awc.ExchangeRate * 1000 / 24 / 60 / 60
		freshSpeed  = mlPerSecond * float64(j.freshCalibration.Steps) / j.freshCalibration.Volume
		wasteSpeed  = mlPerSecond * float64(j.wasteCalibration.Steps) / j.wasteCalibration.Volume
		// freshSteps           = freshSpeed * targetDurationSec * 2
		// wasteSteps           = wasteSpeed * targetDurationSec * 2
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
				// } else {
				// 	var (
				// 		totalFreshSteps  = currentFreshPosition - initialFreshPosition
				// 		totalFreshVolume = totalFreshSteps * freshCalibration.MeasuredVolume / freshCalibration.TargetVolume
				// 	)
			}

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

			// NOTE: Let's adjust the speed every time and re-calculate the target position based on time since intitial
			// err = j.freshFirmata.StepperTo(int(j.freshPump.DeviceID), currentFreshPosition.Position+freshSteps)
			// if err != nil {
			// 	j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("stepping fresh pump (aborting job run): %w", err)}
			// 	continue
			// }
			// err = j.wasteFirmata.StepperTo(int(j.wastePump.DeviceID), currentFreshPosition.Position+wasteSteps)
			// if err != nil {
			// 	j.eventCh <- &AWCJobError{j.awc, fmt.Errorf("stepping waste pump (aborting job run): %w", err)}
			// 	continue
			// }
			//
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
