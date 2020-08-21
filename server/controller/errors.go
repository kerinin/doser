package controller

import (
	"fmt"

	"github.com/kerinin/doser/service/models"
)

type ATOJobError struct {
	AutoTopOff *models.AutoTopOff
	Err        error
}

func (e *ATOJobError) Error() string {
	return e.Err.Error()
}

func (e *ATOJobError) Unwrap() error { return e.Err }

type MaxFillVolumeError struct {
	AutoTopOff *models.AutoTopOff
}

func (e *MaxFillVolumeError) Error() string {
	return fmt.Sprintf("reached max fill volume for auto top off %s", e.AutoTopOff.ID)
}

type UncontrolledPumpError struct {
	PumpID string
	Err    error
}

func (e *UncontrolledPumpError) Unwrap() error { return e.Err }

func (e *UncontrolledPumpError) Error() string {
	return fmt.Sprintf("pump %s left in uncontrolled state: %w", e.PumpID, e.Err)
}
