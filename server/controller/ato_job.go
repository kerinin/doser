package controller

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gobot.io/x/gobot/platforms/raspi"
)

type ATOJob struct {
	controller  *ATO
	ato         *models.AutoTopOff
	pump        *models.Pump
	firmata     *gomata.Firmata
	sensors     []*models.WaterLevelSensor
	calibration *models.Calibration
	mx          sync.Mutex
}

func NewATOJob(
	controller *ATO,
	ato *models.AutoTopOff,
	pump *models.Pump,
	firmata *gomata.Firmata,
	sensors []*models.WaterLevelSensor,
	calibration *models.Calibration,
) *ATOJob {
	return &ATOJob{
		controller:  controller,
		ato:         ato,
		pump:        pump,
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

	interval := time.Duration(j.ato.FillInterval) * time.Minute
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Run first job immediately
	jobCtx, _ := context.WithTimeout(ctx, interval)
	j.runJob(jobCtx, maxSteps, speed)

	for {
		select {
		case <-ticker.C:
			jobCtx, _ := context.WithTimeout(ctx, interval)
			j.runJob(jobCtx, maxSteps, speed)
		case <-ctx.Done():
			return
		}
	}
}

func (j *ATOJob) event(kind string, data string, args ...interface{}) {
	j.controller.eventCh <- &models.AtoEvent{
		ID:           uuid.New().String(),
		AutoTopOffID: j.ato.ID,
		Timestamp:    time.Now().Unix(),
		Kind:         kind,
		Data:         fmt.Sprintf(data, args...),
	}
}

func (j *ATOJob) runJob(ctx context.Context, maxSteps, speed int32) {
	// Connect to the RPi
	rpi := raspi.NewAdaptor()
	err := rpi.Connect()
	if err != nil {
		j.event(ATOJobErrorKind, "Failure connecting to  RPi (aborting job run): %w", err)
		return
	}

	// Ensure the water level sensors are functioning and water isn't currently detected
	for _, sensor := range j.sensors {
		// Read the sensor's current value
		detected, err := WaterDetected(ctx, rpi, j.controller.firmatas, sensor)
		if err != nil {
			j.event(ATOJobErrorKind, "Failure reading water level sensor: %s", err)
			return
		}
		if detected {
			j.event(ATOJobCompleteKind, "Completed ATO - detected water at beginning of job")
			return
		}
	}

	err = j.firmata.StepperZero(int(j.pump.DeviceID))
	if err != nil {
		j.event(ATOJobErrorKind, "Failure zeroing stepper (aborting job run): %w", err)
		return
	}

	if j.pump.Acceleration.Valid {
		err = j.firmata.StepperSetAcceleration(int(j.pump.DeviceID), float32(j.pump.Acceleration.Float64))
		if err != nil {
			j.event(ATOJobErrorKind, "Failure setting pump speed: %w", err)
			return
		}
	}

	err = j.firmata.StepperSetSpeed(int(j.pump.DeviceID), float32(speed))
	if err != nil {
		j.event(ATOJobErrorKind, "Failure setting pump speed (aborting job run): %w", err)
		return
	}

	// Command the stepper to pump the maximum fill volume (we'll interrupt it when a sensor is triggered)
	err = j.firmata.StepperStep(int(j.pump.DeviceID), int32(maxSteps))
	if err != nil {
		j.event(ATOJobErrorKind, "Failure stepping pump (aborting job run): %w", err)
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
				detected, err := WaterDetected(ctx, rpi, j.controller.firmatas, sensor)
				if err != nil {
					j.event(ATOJobErrorKind, "Failure reading water level sensor: %w", err)
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
					j.event(UncontrolledPumpKind, "Failure stopping pump %s after sensor detected water: %w", j.pump.ID, err)
					return
				}

				// if water is detected for alert sensors make some noise
				if sensor.Kind == string(model.SensorKindAlert) {
					j.event(WaterLevelAlertKind, "Water detected for sensor %s of kind %s", sensor.ID, sensor.Kind)
					return
				}

				select {
				case report := <-reportCh:
					// TODO: Convert position to volume
					duration := time.Now().Sub(startTime)
					volume := float64(report.Position) * j.calibration.Volume / float64(j.calibration.Steps)
					j.recordDose(ctx, volume, "ATO %s: %fs", j.ato.ID, duration.Seconds())
					j.event(ATOJobCompleteKind, "Completed ATO - filled %fmL in %fs", volume, duration.Seconds())
					return
				case <-ctx.Done():
					return
				}
			}

		case <-complete:
			// We reached the max fill volume, make some noise
			j.event(MaxFillVolumeErrorKind, "Reached maximum fill volume without detecting water")
			j.recordDose(ctx, j.ato.MaxFillVolume, "ATO %s: max-fill", j.ato.ID)
			return

		case <-ctx.Done():
			// If we timed out, reconnect to firmata and recreate the ATO jobs
			if ctx.Err() == context.DeadlineExceeded {
				err = j.controller.firmatas.Reset()
				if err != nil {
					log.Printf("Failed to reset firmatas: %w", err)
				}
				// Give the firmata a second to clear the serial connection
				<-time.After(time.Second)
				j.controller.Reset()
			}

			log.Printf("Job context cancelled, terminating")

			err = j.firmata.StepperStop(int(j.pump.DeviceID))
			if err != nil {
				j.event(UncontrolledPumpKind, "Failure stopping pump %s during shutdown of ATO job: %w", j.pump.ID, err)
			}
			return
		}
	}
}

func (j *ATOJob) recordDose(ctx context.Context, volume float64, message string, args ...interface{}) {
	dose := models.Dose{
		ID:        uuid.New().String(),
		PumpID:    j.pump.ID,
		Timestamp: time.Now().Unix(),
		Volume:    volume,
		Message:   null.StringFrom(fmt.Sprintf(message, args...)),
	}
	err := dose.Insert(ctx, j.controller.db, boil.Infer())
	if err != nil {
		j.event(ATOJobErrorKind, "Failure to insert dose: %w", err)
	}
}
