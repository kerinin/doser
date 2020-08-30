package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/controller"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	cron "github.com/robfig/cron/v3"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (r *mutationResolver) CreateFirmata(ctx context.Context, serialPort string, baud int) (*models.Firmata, error) {
	m := &models.Firmata{
		ID:         uuid.New().String(),
		SerialPort: serialPort,
		Baud:       int64(baud),
	}
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting firmata: %w", err)
	}

	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) DeleteFirmata(ctx context.Context, id string) (bool, error) {
	f := &models.Firmata{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting firmata: %w", err)
	}

	r.firmatasController.Reset()

	return rows > 0, nil
}

func (r *mutationResolver) CreatePump(ctx context.Context, firmataID string, deviceID string, stepPin int, dirPin *int, enPin *int) (*models.Pump, error) {
	m := &models.Pump{
		ID:        uuid.New().String(),
		FirmataID: firmataID,
		StepPin:   int64(stepPin),
	}
	if dirPin != nil {
		m.DirPin = null.Int64From(int64(*dirPin))
	}
	if enPin != nil {
		m.EnPin = null.Int64From(int64(*enPin))
	}

	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting pump: %w", err)
	}

	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) DeletePump(ctx context.Context, id string) (bool, error) {
	f := &models.Pump{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting pump: %w", err)
	}

	r.firmatasController.Reset()

	return rows > 0, nil
}

func (r *mutationResolver) CalibratePump(ctx context.Context, pumpID string, steps int, volume float64) (*models.Calibration, error) {
	m := &models.Calibration{
		ID:     uuid.New().String(),
		PumpID: pumpID,
		Steps:  int64(steps),
		Volume: volume,
	}
	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting calibration: %w", err)
	}

	r.atoController.Reset()
	r.awcController.Reset()
	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) CreateWaterLevelSensor(ctx context.Context, pin int, kind model.SensorKind, firmataID *string, detectionThreshold *int) (*models.WaterLevelSensor, error) {
	m := &models.WaterLevelSensor{
		ID:        uuid.New().String(),
		Pin:       int64(pin),
		Kind:      kind.String(),
		FirmataID: null.StringFromPtr(firmataID),
	}
	if detectionThreshold != nil {
		m.DetectionThreshold = null.Int64From(int64(*detectionThreshold))
	}

	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting water level sensor: %w", err)
	}

	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) UpdateWaterLevelSensor(ctx context.Context, id string, pin int, kind model.SensorKind, firmataID *string, detectionThreshold *int) (*models.WaterLevelSensor, error) {
	m := &models.WaterLevelSensor{
		ID:        id,
		Pin:       int64(pin),
		Kind:      kind.String(),
		FirmataID: null.StringFromPtr(firmataID),
	}
	if detectionThreshold != nil {
		m.DetectionThreshold = null.Int64From(int64(*detectionThreshold))
	}

	_, err := m.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting water level sensor: %w", err)
	}

	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) DeleteWaterLevelSensor(ctx context.Context, id string) (bool, error) {
	f := &models.WaterLevelSensor{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting water level sensor: %w", err)
	}

	r.firmatasController.Reset()

	return rows > 0, nil
}

func (r *mutationResolver) CreateAutoTopOff(ctx context.Context, pumpID string, levelSensors []string, fillRate float64, fillFrequency string, maxFillVolume float64) (*models.AutoTopOff, error) {
	_, err := cron.ParseStandard(fillFrequency)
	if err != nil {
		return nil, fmt.Errorf("parsing fill frequency as cron: %w", err)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
		err = tx.Commit()
	}()

	pump, err := models.FindPump(ctx, tx, pumpID)
	if err != nil {
		return nil, fmt.Errorf("finding ATO pump: %w", err)
	}

	_, err = pump.Calibrations().One(ctx, tx)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("refusing to create ATO with uncalibrated pump")
	} else if err != nil {
		return nil, fmt.Errorf("getting pump calibration: %w", err)
	}

	m := &models.AutoTopOff{
		ID:            uuid.New().String(),
		PumpID:        pumpID,
		FillRate:      fillRate,
		FillFrequency: fillFrequency,
		MaxFillVolume: maxFillVolume,
	}
	waterLevelSensors := make([]*models.WaterLevelSensor, 0, len(levelSensors))
	for _, sensor := range levelSensors {
		waterLevelSensors = append(waterLevelSensors, &models.WaterLevelSensor{ID: sensor})
	}

	err = m.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting auto top off: %w", err)
	}

	err = m.SetWaterLevelSensors(ctx, tx, false, waterLevelSensors...)
	if err != nil {
		return nil, fmt.Errorf("associating water level sensor: %w")
	}

	r.atoController.Reset()
	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) DeleteAutoTopOff(ctx context.Context, id string) (bool, error) {
	f := &models.AutoTopOff{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting auto top off: %w", err)
	}

	r.atoController.Reset()
	r.firmatasController.Reset()

	return rows > 0, nil
}

func (r *mutationResolver) CreateAutoWaterChange(ctx context.Context, freshPumpID string, wastePumpID string, exchangeRate float64) (*models.AutoWaterChange, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
		err = tx.Commit()
	}()

	freshPump, err := models.FindPump(ctx, tx, freshPumpID)
	if err != nil {
		return nil, fmt.Errorf("finding fresh pump: %w", err)
	}

	_, err = freshPump.Calibrations().One(ctx, tx)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("refusing to create ATO with uncalibrated fresh pump")
	} else if err != nil {
		return nil, fmt.Errorf("getting pump calibration: %w", err)
	}

	wastePump, err := models.FindPump(ctx, tx, wastePumpID)
	if err != nil {
		return nil, fmt.Errorf("finding waste pump: %w", err)
	}

	_, err = wastePump.Calibrations().One(ctx, tx)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("refusing to create ATO with uncalibrated waste pump")
	} else if err != nil {
		return nil, fmt.Errorf("getting pump calibration: %w", err)
	}

	m := &models.AutoWaterChange{
		ID:           uuid.New().String(),
		FreshPumpID:  freshPumpID,
		WastePumpID:  wastePumpID,
		ExchangeRate: exchangeRate,
	}

	err = m.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting auto water change: %w", err)
	}

	r.awcController.Reset()

	return m, nil
}

func (r *mutationResolver) DeleteAutoWaterChange(ctx context.Context, id string) (bool, error) {
	f := &models.AutoWaterChange{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting auto water change: %w", err)
	}

	r.awcController.Reset()

	return rows > 0, nil
}

func (r *mutationResolver) CreateDoser(ctx context.Context, input model.DoserInput) (*models.Doser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDoser(ctx context.Context, id string) (bool, error) {
	f := &models.Doser{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting doser: %w", err)
	}

	return rows > 0, nil
}

func (r *mutationResolver) Pump(ctx context.Context, pumpID string, steps int, speed float64) (bool, error) {
	pump, err := models.FindPump(ctx, r.db, pumpID)
	if err != nil {
		return false, fmt.Errorf("finding pump: %w", err)
	}

	firmata, err := r.firmatasController.Get(ctx, pump.FirmataID)
	if err != nil {
		return false, fmt.Errorf("getting pump's firmata: %w", err)
	}

	err = controller.ConfigurePump(pump, firmata)
	if err != nil {
		return false, fmt.Errorf("configuring pump: %w", err)
	}

	err = firmata.StepperSetSpeed(int(pump.DeviceID), float32(speed))
	if err != nil {
		return false, fmt.Errorf("setting stepper speed: %w", err)
	}

	err = firmata.StepperStep(int(pump.DeviceID), int32(steps))
	if err != nil {
		return false, fmt.Errorf("starting stepper: %w", err)
	}

	complete := firmata.AwaitStepperMoveCompletion(int32(pump.DeviceID))

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-complete:
		return true, nil
	}
}

func (r *pumpResolver) Firmata(ctx context.Context, obj *models.Pump) (*models.Firmata, error) {
	firmata, err := obj.Firmatum().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting firmata: %w", err)
	}

	err = r.firmatasController.Reset()
	if err != nil {
		return nil, fmt.Errorf("resetting firmatas: %w", err)
	}
	r.awcController.Reset()
	r.atoController.Reset()

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

func (r *waterLevelSensorResolver) FirmataID(ctx context.Context, obj *models.WaterLevelSensor) (*string, error) {
	if obj.FirmataID.Valid {
		return &obj.FirmataID.String, nil
	}
	return nil, nil
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
	return controller.WaterDetected(ctx, r.firmatasController, obj)
}

func (r *waterLevelSensorResolver) DetectionThreshold(ctx context.Context, obj *models.WaterLevelSensor) (*int, error) {
	if obj.DetectionThreshold.Valid {
		v := int(obj.DetectionThreshold.Int64)
		return &v, nil
	}
	return nil, nil
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
