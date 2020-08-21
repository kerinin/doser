package controller

import (
	"context"
	"log"

	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/sysfs"
)

type ATOJob struct {
	ctx     context.Context
	firmata *gomata.Firmata
	ato     *models.AutoTopOff
	pump    *models.Pump
	sensors []*models.WaterLevelSensor
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
	rpi := raspi.NewAdaptor()
	err := rpi.Connect()
	if err != nil {
		log.Printf("Unable to connect to RPi: %s", err)
	}

	// TODO: Calibration...

	j.firmata.StepperSetSpeed(int(j.pump.DeviceID), float32(j.ato.FillRate))
	j.firmata.StepperStep(int(j.pump.DeviceID), int32(j.ato.MaxFillVolume))

	// complete := j.firmata.AwaitMoveCompletion(j.pump.DeviceID)

	for {
		select {
		case <-j.ctx.Done():
			return
		// case <-complete:
		// 	// We reached the max fill volume, make some noise
		// 	return
		default:
			for _, sensor := range j.sensors {
				// Read the sensor's current value
				val, err := rpi.DigitalRead(string(sensor.Pin))
				if err != nil {
					log.Printf("Unable to read from RPi: %s", err)
					j.firmata.StepperStop(int(j.pump.DeviceID))

					// TODO: make noise
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
		}
	}
}
