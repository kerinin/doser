package controller

import (
	"fmt"

	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
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
