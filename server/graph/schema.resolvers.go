package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/sysfs"
)

func (r *autoTopOffResolver) Pump(ctx context.Context, obj *models.AutoTopOff) (*models.Pump, error) {
	m, err := obj.Pump().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting ATO: %w", err)
	}

	return m, nil
}

func (r *autoTopOffResolver) LevelSensors(ctx context.Context, obj *models.AutoTopOff) ([]*models.WaterLevelSensor, error) {
	ms, err := obj.WaterLevelSensors().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting level sensors: %w", err)
	}

	return ms, nil
}

func (r *autoWaterChangeResolver) FreshPump(ctx context.Context, obj *models.AutoWaterChange) (*models.Pump, error) {
	m, err := obj.FreshPump().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting fresh water pump: %w", err)
	}

	return m, nil
}

func (r *autoWaterChangeResolver) WastePump(ctx context.Context, obj *models.AutoWaterChange) (*models.Pump, error) {
	m, err := obj.WastePump().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting waste water pump: %w", err)
	}

	return m, nil
}

func (r *doserResolver) Components(ctx context.Context, obj *models.Doser) ([]*models.DoserComponent, error) {
	ms, err := obj.DoserComponents().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting doser components: %w", err)
	}

	return ms, nil
}

func (r *doserComponentResolver) Pump(ctx context.Context, obj *models.DoserComponent) (*models.Pump, error) {
	m, err := obj.Pump().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting pump: %w", err)
	}

	return m, nil
}

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

func (r *mutationResolver) CreateWaterLevelSensor(ctx context.Context, input model.CreateWaterLevelSensor) (*models.WaterLevelSensor, error) {
	m := &models.WaterLevelSensor{
		ID:   uuid.New().String(),
		Pin:  int64(input.Pin),
		Kind: input.Kind.String(),
	}

	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting water level sensor: %w", err)
	}

	return m, nil
}

func (r *mutationResolver) CreateAutoTopOff(ctx context.Context, input model.NewAutoTopOff) (*models.AutoTopOff, error) {
	// NOTE: Be sure to parse the fill frequency to verify it's a valid cron
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAutoWaterChange(ctx context.Context, input model.NewAutoWaterChangeInput) (*models.AutoWaterChange, error) {
	m := &models.AutoWaterChange{
		ID:           uuid.New().String(),
		FreshPumpID:  input.FreshPumpID,
		WastePumpID:  input.WastePumpID,
		ExchangeRate: input.ExchangeRate,
	}

	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting auto water change: %w", err)
	}

	return m, nil
}

func (r *mutationResolver) CreateDoser(ctx context.Context, input model.NewDoserInput) (*models.Doser, error) {
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
	return NullInt64ToIntPtr(obj.EnPin), nil
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

func (r *queryResolver) WaterLevelSensors(ctx context.Context) ([]*models.WaterLevelSensor, error) {
	ms, err := models.WaterLevelSensors().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting water level sensors: %w", err)
	}

	return ms, nil
}

func (r *queryResolver) AutoTopOff(ctx context.Context) ([]*models.AutoTopOff, error) {
	ms, err := models.AutoTopOffs().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting ATOs: %w", err)
	}

	return ms, nil
}

func (r *queryResolver) AutoWaterChanges(ctx context.Context) ([]*models.AutoWaterChange, error) {
	ms, err := models.AutoWaterChanges().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting AWCs: %w", err)
	}

	return ms, nil
}

func (r *queryResolver) Dosers(ctx context.Context) ([]*models.Doser, error) {
	ms, err := models.Dosers().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting Dosers: %w", err)
	}

	return ms, nil
}

func (r *waterLevelSensorResolver) Kind(ctx context.Context, obj *models.WaterLevelSensor) (model.SensorKind, error) {
	for _, kind := range model.AllSensorKind {
		if string(kind) == obj.Kind {
			return kind, nil
		}
	}
	return model.SensorKind(""), fmt.Errorf("Unrecognized sensor kind %s", obj.Kind)
}

func (r *waterLevelSensorResolver) WaterDetected(ctx context.Context, obj *models.WaterLevelSensor) (bool, error) {
	rpi := raspi.NewAdaptor()
	err := rpi.Connect()
	if err != nil {
		return false, fmt.Errorf("connecting to rpi: %w", err)
	}

	pinString := strconv.Itoa(int(obj.Pin))
	val, err := rpi.DigitalRead(pinString)
	if err != nil {
		return false, fmt.Errorf("reading sensor pin %s: %w", pinString, err)
	}

	return val == sysfs.HIGH, nil
}

// AutoTopOff returns generated.AutoTopOffResolver implementation.
func (r *Resolver) AutoTopOff() generated.AutoTopOffResolver { return &autoTopOffResolver{r} }

// AutoWaterChange returns generated.AutoWaterChangeResolver implementation.
func (r *Resolver) AutoWaterChange() generated.AutoWaterChangeResolver {
	return &autoWaterChangeResolver{r}
}

// Doser returns generated.DoserResolver implementation.
func (r *Resolver) Doser() generated.DoserResolver { return &doserResolver{r} }

// DoserComponent returns generated.DoserComponentResolver implementation.
func (r *Resolver) DoserComponent() generated.DoserComponentResolver {
	return &doserComponentResolver{r}
}

// Firmata returns generated.FirmataResolver implementation.
func (r *Resolver) Firmata() generated.FirmataResolver { return &firmataResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Pump returns generated.PumpResolver implementation.
func (r *Resolver) Pump() generated.PumpResolver { return &pumpResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// WaterLevelSensor returns generated.WaterLevelSensorResolver implementation.
func (r *Resolver) WaterLevelSensor() generated.WaterLevelSensorResolver {
	return &waterLevelSensorResolver{r}
}

type autoTopOffResolver struct{ *Resolver }
type autoWaterChangeResolver struct{ *Resolver }
type doserResolver struct{ *Resolver }
type doserComponentResolver struct{ *Resolver }
type firmataResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type pumpResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type waterLevelSensorResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *waterLevelSensorResolver) Firmata(ctx context.Context, obj *models.WaterLevelSensor) (*models.Firmata, error) {
	m, err := obj.Firmatum().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting firmata: %w", err)
	}

	return m, nil
}
func (r *pumpResolver) DeviceID(ctx context.Context, obj *models.Pump) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
