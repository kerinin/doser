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
	ctx     context.Context
	db      *sql.DB
	firmata *gomata.Firmata
	ato     *models.AutoTopOff
	pump    *models.Pump
	sensors []*models.WaterLevelSensor
	errs    chan error
	mx      sync.Mutex
}

func NewATOJob(ato *models.AutoTopOff) *ATOJob {
	return &ATOJob{
		ctx:     context.TODO(),
		firmata: nil,
		ato:     ato,
		pump:    nil,
		sensors: nil,
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

	// TODO: Calibration...
	calibration, err := j.pump.Calibrations().One(j.ctx, j.db)
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

	err = j.firmata.StepperSetSpeed(int(j.pump.DeviceID), float32(speed))
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("setting pump speed (aborting job run): %w", err)}
		return
	}

	err = j.firmata.StepperStep(int(j.pump.DeviceID), int32(maxSteps))
	if err != nil {
		j.errs <- &ATOJobError{j.ato, fmt.Errorf("stepping pump (aborting job run): %w", err)}
		return
	}

	complete := j.firmata.AwaitStepperMoveCompletion(int32(j.pump.DeviceID))
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, sensor := range j.sensors {
				// Read the sensor's current value
				val, err := rpi.DigitalRead(string(sensor.Pin))
				if err != nil {
					j.errs <- &ATOJobError{j.ato, fmt.Errorf("connecting to RPi: %w", err)}
					if err := j.firmata.StepperStop(int(j.pump.DeviceID)); err != nil {
						j.errs <- &UncontrolledPumpError{j.pump.ID, fmt.Errorf("stopping pump after failing to read sensor: %w", err)}
					}
					return
				}

				if val == sysfs.LOW {
					// No water detected, keep checking
					continue
				}

				// water detected, stop stepper
				j.firmata.StepperStop(int(j.pump.DeviceID))

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
			err = j.firmata.StepperStop(int(j.pump.DeviceID))
			if err != nil {
				j.errs <- &UncontrolledPumpError{j.pump.ID, fmt.Errorf("stopping pump during shutdown of ATO job: %w", err)}
			}
			return
		}
	}
}
