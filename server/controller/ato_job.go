package controller

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/sysfs"
)

type ATOJob struct {
	ctx         context.Context
	wg          *sync.WaitGroup
	eventCh     chan<- Event
	ato         *models.AutoTopOff
	pump        *models.Pump
	firmata     *gomata.Firmata
	sensors     []*models.WaterLevelSensor
	calibration *models.Calibration
	mx          sync.Mutex
}

func NewATOJob(
	ctx context.Context,
	wg *sync.WaitGroup,
	eventCh chan<- Event,
	ato *models.AutoTopOff,
	pump *models.Pump,
	firmata *gomata.Firmata,
	sensors []*models.WaterLevelSensor,
	calibration *models.Calibration,
) *ATOJob {
	return &ATOJob{
		ctx:         ctx,
		wg:          wg,
		eventCh:     eventCh,
		ato:         ato,
		pump:        pump,
		firmata:     firmata,
		sensors:     sensors,
		calibration: calibration,
	}
}

func (j *ATOJob) Run() {
	// Prevent multiple jobs from running concurrently
	j.mx.Lock()
	defer j.mx.Unlock()

	log.Printf("Running ATO job %s", j.ato.ID)

	// Connect to the RPi
	rpi := raspi.NewAdaptor()
	err := rpi.Connect()
	if err != nil {
		j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("connecting to  RPi (aborting job run): %w", err)}
		return
	}

	// Ensure the water level sensors are functioning and water isn't currently detected
	for _, sensor := range j.sensors {
		// Read the sensor's current value
		val, err := rpi.DigitalRead(string(sensor.Pin))
		if err != nil {
			j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("connecting to RPi: %w", err)}
			return
		}
		if val == sysfs.HIGH {
			j.eventCh <- &ATOJobComplete{j.ato, 0}
			return
		}
	}

	// Configure the stepper
	var (
		maxSteps = j.ato.MaxFillVolume * float64(j.calibration.Steps) / j.calibration.Volume
		speed    = j.ato.FillRate * float64(j.calibration.Steps) / j.calibration.Volume
	)
	err = j.firmata.StepperSetSpeed(int(j.pump.DeviceID), float32(speed))
	if err != nil {
		j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("setting pump speed (aborting job run): %w", err)}
		return
	}

	// Prevent termination during the run
	j.wg.Add(1)
	defer j.wg.Done()

	// Command the stepper to pump the maximum fill volume (we'll interrupt it when a sensor is triggered)
	err = j.firmata.StepperStep(int(j.pump.DeviceID), int32(maxSteps))
	if err != nil {
		j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("stepping pump (aborting job run): %w", err)}
		return
	}

	var (
		startTime = time.Now()
		complete  = j.firmata.AwaitStepperMoveCompletion(int32(j.pump.DeviceID))
		ticker    = time.NewTicker(10 * time.Millisecond)
	)
	defer ticker.Stop()

	// Wait for completion (or termination)
	for {
		select {
		case <-ticker.C:
			for _, sensor := range j.sensors {
				// Read the sensor's current value
				val, err := rpi.DigitalRead(string(sensor.Pin))
				if err != nil {
					j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("connecting to RPi: %w", err)}
					if err := j.firmata.StepperStop(int(j.pump.DeviceID)); err != nil {
						j.eventCh <- &UncontrolledPumpError{j.pump.ID, fmt.Errorf("stopping pump after failing to read sensor: %w", err)}
					}
					return
				}

				if val == sysfs.LOW {
					// No water detected, keep checking
					continue
				}

				// water detected, stop stepper
				err = j.firmata.StepperStop(int(j.pump.DeviceID))
				if err != nil {
					j.eventCh <- &UncontrolledPumpError{j.pump.ID, fmt.Errorf("stopping pump after sensor detected water: %w", err)}
					return
				}

				// if water is detected for alert sensors make some noise
				if sensor.Kind == string(model.SensorKindAlert) {
					j.eventCh <- &WaterLevelAlert{sensor}
					return
				}

				duration := time.Now().Sub(startTime)
				j.eventCh <- &ATOJobComplete{j.ato, duration}
				return
			}

		case <-complete:
			// We reached the max fill volume, make some noise
			j.eventCh <- &MaxFillVolumeError{j.ato}
			return

		case <-j.ctx.Done():
			err = j.firmata.StepperStop(int(j.pump.DeviceID))
			if err != nil {
				j.eventCh <- &UncontrolledPumpError{j.pump.ID, fmt.Errorf("stopping pump during shutdown of ATO job: %w", err)}
			}
			return
		}
	}
}
