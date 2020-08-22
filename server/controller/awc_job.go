package controller

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

const targetDurationSec = 60

type AWCJob struct {
	db       *sql.DB
	firmatas map[string]*gomata.Firmata
	awc      *models.AutoWaterChange
}

func NewAWCJob(db *sql.DB, firmatas map[string]*gomata.Firmata, awc *models.AutoWaterChange) *AWCJob {
	return &AWCJob{db, firmatas, awc}
}

func (j *AWCJob) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (j *AWCJob) runOnce(ctx context.Context, wg *sync.WaitGroup) Event {
	// Fetch resources necessary for the ATO job
	freshPump, err := j.awc.FreshPump().One(ctx, j.db)
	if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("getting fresh water pump (aborting job run): %w", err)}
	}
	wastePump, err := j.awc.WastePump().One(ctx, j.db)
	if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("getting waste water pump (aborting job run): %w", err)}
	}
	freshFirmata, found := j.firmatas[freshPump.FirmataID]
	if !found {
		return &AWCJobError{j.awc, fmt.Errorf("unrecognized firmata ID %s for fresh pump %s (aborting job run)", freshPump.FirmataID, freshPump.ID)}
	}
	wasteFirmata, found := j.firmatas[wastePump.FirmataID]
	if !found {
		return &AWCJobError{j.awc, fmt.Errorf("unrecognized firmata ID %s for waste pump %s (aborting job run)", wastePump.FirmataID, wastePump.ID)}
	}
	freshCalibration, err := freshPump.Calibrations().One(ctx, j.db)
	if err == sql.ErrNoRows {
		return &AWCJobError{j.awc, fmt.Errorf("refusing to run ATO job with uncalibrated fresh pump")}
	} else if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("getting pump calibration (aborting job run): %w", err)}
	}
	wasteCalibration, err := wastePump.Calibrations().One(ctx, j.db)
	if err == sql.ErrNoRows {
		return &AWCJobError{j.awc, fmt.Errorf("refusing to run ATO job with uncalibrated waste pump")}
	} else if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("getting pump calibration (aborting job run): %w", err)}
	}

	// Configure the stepper
	// Exchange rate is in L/day
	var (
		mlPerSecond = j.awc.ExchangeRate / 1000 / 24 / 60 / 60
		freshSpeed  = mlPerSecond * freshCalibration.TargetVolume / freshCalibration.MeasuredVolume
		wasteSpeed  = mlPerSecond * wasteCalibration.TargetVolume / wasteCalibration.MeasuredVolume
		freshSteps  = freshSpeed * targetDurationSec
		wasteSteps  = wasteSpeed * targetDurationSec
	)
	err = freshFirmata.StepperSetSpeed(int(freshPump.DeviceID), float32(freshSpeed))
	if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("setting fresh pump speed (aborting job run): %w", err)}
	}
	err = wasteFirmata.StepperSetSpeed(int(wastePump.DeviceID), float32(wasteSpeed))
	if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("setting waste pump speed (aborting job run): %w", err)}
	}

	// Prevent termination during the run
	wg.Add(1)
	defer wg.Done()

	// Command the stepper to pump the maximum fill volume (we'll interrupt it when a sensor is triggered)
	err = freshFirmata.StepperStep(int(freshPump.DeviceID), int32(freshSteps))
	if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("stepping fresh pump (aborting job run): %w", err)}
	}
	err = wasteFirmata.StepperStep(int(wastePump.DeviceID), int32(wasteSteps))
	if err != nil {
		return &AWCJobError{j.awc, fmt.Errorf("stepping waste pump (aborting job run): %w", err)}
	}

	var (
		startTime     = time.Now()
		freshComplete = freshFirmata.AwaitStepperMoveCompletion(int32(freshPump.DeviceID))
		wasteComplete = wasteFirmata.AwaitStepperMoveCompletion(int32(wastePump.DeviceID))
		ticker        = time.NewTicker(10 * time.Millisecond)
	)
	defer ticker.Stop()

	// Wait for completion (or termination)
	for {
		select {
		// TODO: Wait for both pumps to complete...
		case <-freshComplete:
			duration := time.Now().Sub(startTime)
			return &AWCJobComplete{j.awc, duration}

		case <-wasteComplete:
			duration := time.Now().Sub(startTime)
			return &AWCJobComplete{j.awc, duration}

		case <-ctx.Done():
			// TODO: Allow both of these to fail
			err = freshFirmata.StepperStop(int(freshPump.DeviceID))
			if err != nil {
				return &UncontrolledPumpError{freshPump.ID, fmt.Errorf("stopping fresh pump during shutdown of ATO job: %w", err)}
			}
			err = wasteFirmata.StepperStop(int(wastePump.DeviceID))
			if err != nil {
				return &UncontrolledPumpError{wastePump.ID, fmt.Errorf("stopping waste pump during shutdown of ATO job: %w", err)}
			}
			duration := time.Now().Sub(startTime)
			return &AWCJobComplete{j.awc, duration}
		}
	}
}
