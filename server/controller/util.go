package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/sysfs"
)

func ConfigurePump(pump *models.Pump, firmata *gomata.Firmata) error {
	var (
		deviceID  = int(pump.DeviceID)
		wireCount = gomata.Driver
		stepType  = gomata.WholeStep
		hasEnable = gomata.EnablePin
		pin1      = int(pump.StepPin)
		pin2      = int(pump.DirPin.Int64)
		pin3      = 0
		pin4      = 0
		enablePin = int(pump.EnPin.Int64)
		invert    = gomata.Inversions(0)
	)
	if pump.EnPin.IsZero() {
		hasEnable = gomata.NoEnablePin
	}

	err := firmata.StepperConfigure(deviceID, wireCount, stepType, hasEnable, pin1, pin2, pin3, pin4, enablePin, invert)
	if err != nil {
		return fmt.Errorf("configuring stepper: %w", err)
	}
	return nil
}

func WaterDetected(ctx context.Context, rpi *raspi.Adaptor, firmatasController *Firmatas, obj *models.WaterLevelSensor) (bool, error) {
	if !obj.FirmataID.Valid {
		return gpioWaterDetected(ctx, rpi, obj)
	}

	firmata, err := firmatasController.Get(ctx, obj.FirmataID.String)
	if err != nil {
		return false, fmt.Errorf("getting firmata: %w", err)
	}

	// Pause a bit to let pin values arrive
	<-time.After(10 * time.Millisecond)

	var (
		// Defaults to zero if not set
		threshold = int(obj.DetectionThreshold.Int64)
		pins      = firmata.Pins()
		val       = pins[obj.Pin].Value
	)

	// XOR invert
	return obj.Invert != (val > threshold), nil
}

func gpioWaterDetected(ctx context.Context, rpi *raspi.Adaptor, obj *models.WaterLevelSensor) (bool, error) {
	pinString := strconv.Itoa(int(obj.Pin))
	val, err := rpi.DigitalRead(pinString)
	if err != nil {
		return false, fmt.Errorf("reading sensor pin %s: %w", pinString, err)
	}

	// XOR invert
	return obj.Invert != (val == sysfs.HIGH), nil
}
