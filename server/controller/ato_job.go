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
)

type ATOJob struct {
	eventCh     chan<- Event
	ato         *models.AutoTopOff
	pump        *models.Pump
	firmatas    *Firmatas
	firmata     *gomata.Firmata
	sensors     []*models.WaterLevelSensor
	calibration *models.Calibration
	mx          sync.Mutex
}

func NewATOJob(
	eventCh chan<- Event,
	ato *models.AutoTopOff,
	pump *models.Pump,
	firmatas *Firmatas,
	firmata *gomata.Firmata,
	sensors []*models.WaterLevelSensor,
	calibration *models.Calibration,
) *ATOJob {
	return &ATOJob{
		eventCh:     eventCh,
		ato:         ato,
		pump:        pump,
		firmatas:    firmatas,
		firmata:     firmata,
		sensors:     sensors,
		calibration: calibration,
	}
}

func (j *ATOJob) Run(ctx context.Context, wg *sync.WaitGroup) {
	// Prevent termination during the run
	wg.Add(1)
	defer wg.Done()

	log.Printf("Starting ATO job %s", j.ato.ID)

	// Configure the stepper
	var (
		maxSteps = int32(j.ato.MaxFillVolume * float64(j.calibration.Steps) / j.calibration.Volume)
		speed    = (int32(j.ato.FillRate*float64(time.Second)*float64(j.calibration.Steps)/j.calibration.Volume/float64(time.Minute)) / 100) * 100
	)
	log.Printf("ATO job params - deviceID:%d maxSteps:%d speed:%d", j.pump.DeviceID, maxSteps, speed)

	ticker := time.NewTicker(time.Duration(j.ato.FillInterval) * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			j.runJob(ctx, maxSteps, speed)
		case <-ctx.Done():
			return
		}
	}
}
func (j *ATOJob) runJob(ctx context.Context, maxSteps, speed int32) {
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
		detected, err := WaterDetected(ctx, rpi, j.firmatas, sensor)
		if err != nil {
			j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("reading water level sensor: %w", err)}
			return
		}
		if detected {
			j.eventCh <- &ATOJobComplete{j.ato, 0, 0}
			return
		}
	}

	err = j.firmata.StepperZero(int(j.pump.DeviceID))
	if err != nil {
		j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("zeroing stepper (aborting job run): %w", err)}
		return
	}

	err = j.firmata.StepperSetSpeed(int(j.pump.DeviceID), float32(speed))
	if err != nil {
		j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("setting pump speed (aborting job run): %w", err)}
		return
	}

	// Command the stepper to pump the maximum fill volume (we'll interrupt it when a sensor is triggered)
	err = j.firmata.StepperStep(int(j.pump.DeviceID), int32(maxSteps))
	if err != nil {
		j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("stepping pump (aborting job run): %w", err)}
		return
	}

	var (
		startTime    = time.Now()
		complete     = j.firmata.AwaitStepperMoveCompletion(int32(j.pump.DeviceID))
		sensorTicker = time.NewTicker(10 * time.Millisecond)
	)
	defer sensorTicker.Stop()

	// Wait for completion (or termination)
	for {
		select {
		case <-sensorTicker.C:
			for _, sensor := range j.sensors {
				// Read the sensor's current value
				detected, err := WaterDetected(ctx, rpi, j.firmatas, sensor)
				if err != nil {
					j.eventCh <- &ATOJobError{j.ato, fmt.Errorf("reading water level sensor: %w", err)}
					return
				}
				if !detected {
					// No water detected, keep checking
					continue
				}

				// water detected, stop stepper
				reportCh := j.firmata.AwaitStepperMoveCompletion(int32(j.pump.DeviceID))

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

				select {
				case report := <-reportCh:
					// TODO: Convert position to volume
					duration := time.Now().Sub(startTime)
					volume := float64(report.Position) * j.calibration.Volume / float64(j.calibration.Steps)
					j.eventCh <- &ATOJobComplete{j.ato, duration, volume}
					return
				case <-ctx.Done():
					return
				}
			}

		case <-complete:
			// We reached the max fill volume, make some noise
			j.eventCh <- &MaxFillVolumeError{j.ato}
			return

		case <-ctx.Done():
			err = j.firmata.StepperStop(int(j.pump.DeviceID))
			if err != nil {
				j.eventCh <- &UncontrolledPumpError{j.pump.ID, fmt.Errorf("stopping pump during shutdown of ATO job: %w", err)}
			}
			return
		}
	}
}
