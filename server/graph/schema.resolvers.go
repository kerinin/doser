package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *firmataResolver) Pumps(ctx context.Context, obj *models.Firmata) ([]*models.Pump, error) {
	pumps, err := obj.FirmatumPumps().All(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting pumps: %w", err)
	}

	return pumps, nil
}

func (r *mutationResolver) CreateFirmata(ctx context.Context, input model.NewFirmataInput) (*models.Firmata, error) {
	m := &models.Firmata{
		ID:         uuid.New().String(),
		SerialPort: input.SerialPort,
	}
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting firmata: %w", err)
	}

	return m, nil
}

func (r *mutationResolver) CreatePump(ctx context.Context, input model.NewPumpInput) (*models.Pump, error) {
	m := &models.Pump{
		ID:        uuid.New().String(),
		FirmataID: input.FirmataID,
		StepPin:   int64(input.StepPin),
	}
	if input.DirPin != nil {
		m.EnPin = null.Int64From(int64(*input.EnPin))
	}
	if input.EnPin != nil {
		m.EnPin = null.Int64From(int64(*input.EnPin))
	}

	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting pump: %w", err)
	}

	return m, nil
}

func (r *mutationResolver) CalibratePump(ctx context.Context, input model.CalibratePumpInput) (*models.Calibration, error) {
	m := &models.Calibration{
		ID:             uuid.New().String(),
		PumpID:         input.PumpID,
		TargetVolume:   input.TargetVolume,
		MeasuredVolume: input.MeasuredVolume,
	}
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting calibration: %w", err)
	}

	return m, nil
}

func (r *mutationResolver) CreateWaterLevelSensor(ctx context.Context, input model.CreateWaterLevelSensor) (*model.WaterLevelSensor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAutoTopOff(ctx context.Context, input model.NewAutoTopOff) (*model.AutoTopOff, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAutoWaterChange(ctx context.Context, input model.NewAutoWaterChangeInput) (*model.AutoWaterChange, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDosers(ctx context.Context, input model.NewDosersInput) (*model.Dosers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pumpResolver) Firmata(ctx context.Context, obj *models.Pump) (*models.Firmata, error) {
	firmata, err := obj.Firmatum().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting firmata: %w", err)
	}

	return firmata, nil
}

func (r *pumpResolver) EnPin(ctx context.Context, obj *models.Pump) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pumpResolver) Calibration(ctx context.Context, obj *models.Pump) (*models.Calibration, error) {
	calibration, err := obj.Calibrations().One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting calibrations: %w", err)
	}

	return calibration, nil
}

func (r *queryResolver) Firmatas(ctx context.Context) ([]*models.Firmata, error) {
	firmatas, err := models.Firmatas().All(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting firmatas: %w", err)
	}

	return firmatas, nil
}

func (r *queryResolver) Pumps(ctx context.Context) ([]*models.Pump, error) {
	pumps, err := models.Pumps().All(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting pumps: %w", err)
	}

	return pumps, nil
}

func (r *queryResolver) WaterLevelSensors(ctx context.Context) ([]*model.WaterLevelSensor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AutoTopOff(ctx context.Context) ([]*model.AutoTopOff, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AutoWaterChanges(ctx context.Context) ([]*model.AutoWaterChange, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dosers(ctx context.Context) ([]*model.Dosers, error) {
	panic(fmt.Errorf("not implemented"))
}

// Firmata returns generated.FirmataResolver implementation.
func (r *Resolver) Firmata() generated.FirmataResolver { return &firmataResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Pump returns generated.PumpResolver implementation.
func (r *Resolver) Pump() generated.PumpResolver { return &pumpResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type firmataResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type pumpResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *pumpResolver) StepPin(ctx context.Context, obj *models.Pump) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *twoPointCalibrationResolver) TargetVolume(ctx context.Context, obj *models.Calibration) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *twoPointCalibrationResolver) MeasuredVolume(ctx context.Context, obj *models.Calibration) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

type twoPointCalibrationResolver struct{ *Resolver }

func (r *queryResolver) Firmata(ctx context.Context) (*models.Firmata, error) {
	panic(fmt.Errorf("not implemented"))
}

const (
	firmataPrefix = "firmata"
	pumpPrefix    = "firmata"
)

func FirmataKey(id string) []byte {
	return []byte(strings.Join([]string{firmataPrefix, id}, "/"))
}
func PumpKey(id string) []byte {
	return []byte(strings.Join([]string{pumpPrefix, id}, "/"))
}
