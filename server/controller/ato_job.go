package controller

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/sysfs"
)

type ATOJob struct {
	ctx      context.Context
	db       *sql.DB
	firmatas map[string]*gomata.Firmata
	ato      *models.AutoTopOff
	errs     chan error
	mx       sync.Mutex
}

func NewATOJob(ctx context.Context, db *sql.DB, firmatas map[string]*gomata.Firmata, ato *models.AutoTopOff) *ATOJob {
	return &ATOJob{
		ctx:      ctx,
		db:       db,
		firmatas: firmatas,
		ato:      ato,
		errs:     make(chan error, 1),
	}
}

func (j *ATOJob) Run() {
	// Prevent multiple jobs from running concurrently
	j.mx.Lock()
	defer j.mx.Unlock()

	rpi := raspi.NewAdaptor()
	err := rpi.Connect()
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("connecting to  RPi (aborting job run): %w", err)}
		return
	}

	pump, err := j.ato.Pump().One(j.ctx, j.db)
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("getting pump (aborting job run): %w", err)}
		return
	}

	firmata, found := j.firmatas[pump.FirmataID]
	if !found {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("unrecognized firmata ID %s for pump %s (aborting job run)", pump.FirmataID, pump.ID)}
		return
	}

	sensors, err := j.ato.WaterLevelSensors().All(j.ctx, j.db)
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("getting sensors (aborting job run): %w", err)}
		return
	}

	// TODO: Calibration...
	calibration, err := pump.Calibrations().One(j.ctx, j.db)
	if err == sql.ErrNoRows {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("refusing to run ATO job with uncalibrated pump")}
		return
	} else if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("getting pump calibration (aborting job run): %w", err)}
		return
	}

	var (
		maxSteps = j.ato.MaxFillVolume * calibration.TargetVolume / calibration.MeasuredVolume
		speed    = j.ato.FillRate * calibration.TargetVolume / calibration.MeasuredVolume
	)

	err = firmata.StepperSetSpeed(int(pump.DeviceID), float32(speed))
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("setting pump speed (aborting job run): %w", err)}
		return
	}

	err = firmata.StepperStep(int(pump.DeviceID), int32(maxSteps))
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("stepping pump (aborting job run): %w", err)}
		return
	}

	complete := firmata.AwaitStepperMoveCompletion(int32(pump.DeviceID))
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, sensor := range sensors {
				// Read the sensor's current value
				val, err := rpi.DigitalRead(string(sensor.Pin))
				if err != nil {
					j.errs <- &ATOJobError{j.ato, fmt.Errorf("connecting to RPi: %w", err)}
					if err := firmata.StepperStop(int(pump.DeviceID)); err != nil {
						j.errs <- &UncontrolledPumpError{pump.ID, fmt.Errorf("stopping pump after failing to read sensor: %w", err)}
					}
					return
				}

				if val == sysfs.LOW {
					// No water detected, keep checking
					continue
				}

				// water detected, stop stepper
				firmata.StepperStop(int(pump.DeviceID))

				// if water is detected for alert sensors make some noise
				if sensor.Kind == string(model.SensorKindAlert) {
					// TODO: make noise
				}
			}

		case <-complete:
			// We reached the max fill volume, make some noise
			j.errs <- &MaxFillVolumeError{j.ato}
			return

		case <-j.ctx.Done():
			err = firmata.StepperStop(int(pump.DeviceID))
			if err != nil {
				j.errs <- &UncontrolledPumpError{pump.ID, fmt.Errorf("stopping pump during shutdown of ATO job: %w", err)}
			}
			return
		}
	}
}
