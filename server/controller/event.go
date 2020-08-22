package controller

import (
	"fmt"
	"time"

	"github.com/kerinin/doser/service/models"
)

type Event interface {
	Message() string
}

type ATOJobComplete struct {
	AutoTopOff *models.AutoTopOff
	Duration   time.Duration
}

func (e *ATOJobComplete) Message() string {
	return fmt.Sprintf("ATO job %s completed in %ds", e.AutoTopOff.ID, e.Duration)
}

type ATOJobError struct {
	AutoTopOff *models.AutoTopOff
	Err        error
}

func (e *ATOJobError) Message() string {
	return fmt.Sprintf("Failure duing ATO job %s: %s", e.AutoTopOff.ID, e.Err)
}

type MaxFillVolumeError struct {
	AutoTopOff *models.AutoTopOff
}

func (e *MaxFillVolumeError) Message() string {
	return fmt.Sprintf("Reached max fill volume for auto top off %s", e.AutoTopOff.ID)
}

type UncontrolledPumpError struct {
	PumpID string
	Err    error
}

func (e *UncontrolledPumpError) Message() string {
	return fmt.Sprintf("Pump %s left in uncontrolled state: %s", e.PumpID, e.Err)
}

type WaterLevelAlert struct {
	Sensor *models.WaterLevelSensor
}

func (e *WaterLevelAlert) Message() string {
	return fmt.Sprintf("Water level sensor %s of kind ALERT detected water", e.Sensor.ID)
}
