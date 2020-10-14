package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/kerinin/doser/service/controller"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/kerinin/doser/service/graph/model"
	"github.com/kerinin/doser/service/models"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gobot.io/x/gobot/platforms/raspi"
)

func (r *autoTopOffResolver) Name(ctx context.Context, obj *models.AutoTopOff) (*string, error) {
	if !obj.Name.Valid {
		return nil, nil
	}
	return &obj.Name.String, nil
}

func (r *autoTopOffResolver) Pump(ctx context.Context, obj *models.AutoTopOff) (*models.Pump, error) {
	m, err := obj.Pump().One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting ATO pump: %w", err)
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

func (r *autoTopOffResolver) FillLevel(ctx context.Context, obj *models.AutoTopOff) (*model.FillLevel, error) {
	if !obj.FillLevelTimestamp.Valid || !obj.FillLevelVolume.Valid {
		return nil, nil
	}
	return &model.FillLevel{int(obj.FillLevelTimestamp.Int64), obj.FillLevelVolume.Float64}, nil
}

func (r *autoTopOffResolver) BurnDown(ctx context.Context, obj *models.AutoTopOff) ([]*model.FillLevel, error) {
	if !obj.FillLevelTimestamp.Valid || !obj.FillLevelVolume.Valid {
		return nil, nil
	}

	p := &models.Pump{ID: obj.PumpID}
	doses, err := p.Doses(
		models.DoseWhere.Timestamp.GT(obj.FillLevelTimestamp.Int64),
	).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting pump doses: %w", err)
	}

	fillLevels := make([]*model.FillLevel, 0, len(doses))
	cumulative := float64(0)
	for _, dose := range doses {
		cumulative -= dose.Volume
		fillLevels = append(fillLevels, &model.FillLevel{
			Timestamp: int(dose.Timestamp),
			Volume:    obj.FillLevelVolume.Float64 - cumulative,
		})
	}

	return nil, nil
}

func (r *autoTopOffResolver) Events(ctx context.Context, obj *models.AutoTopOff) ([]*models.AtoEvent, error) {
	events, err := obj.AtoEvents().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting ATO events: %w", err)
	}

	return events, nil
}

func (r *autoTopOffResolver) Rate(ctx context.Context, obj *models.AutoTopOff, window *int) ([]*model.AtoRate, error) {
	pump := &models.Pump{ID: obj.PumpID}
	doses, err := pump.Doses(qm.OrderBy(models.DoseColumns.Timestamp)).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting dose history: %w", err)
	}

	if len(doses) < 1 {
		return nil, nil
	}

	var (
		windowDuration = int64(60 * 60)
		cursor         = 0
	)
	if window != nil {
		windowDuration = int64(*window)
	}
	var (
		rates     = make([]*model.AtoRate, 0, (doses[len(doses)-1].Timestamp-doses[0].Timestamp)/windowDuration)
		startTime = (doses[0].Timestamp / int64(windowDuration)) * int64(windowDuration)
		endTime   = doses[len(doses)-1].Timestamp
	)

	// Iterate over non-overlapping time windows
	for windowStart := startTime; windowStart < endTime; windowStart += windowDuration {
		var (
			volume    = 0.0
			windowEnd = windowStart + windowDuration
		)

		// Advance the cursor to the start of the current window
		for doses[cursor].Timestamp <= windowStart {
			cursor++
		}
		// Read all the doses in the window
		for i := cursor; i < len(doses) && doses[i].Timestamp <= windowEnd; i++ {
			if doses[i].Volume < 0 {
				continue
			}
			volume += doses[i].Volume
		}

		rate := &model.AtoRate{
			Timestamp: int(windowEnd),
			Rate:      volume * float64(60*60) / float64(windowDuration),
		}
		rates = append(rates, rate)
	}

	return rates, nil
}

func (r *autoWaterChangeResolver) Name(ctx context.Context, obj *models.AutoWaterChange) (*string, error) {
	if !obj.Name.Valid {
		return nil, nil
	}
	return &obj.Name.String, nil
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

func (r *autoWaterChangeResolver) FillLevel(ctx context.Context, obj *models.AutoWaterChange) (*model.FillLevel, error) {
	if !obj.FillLevelTimestamp.Valid || !obj.FillLevelVolume.Valid {
		return nil, nil
	}
	return &model.FillLevel{int(obj.FillLevelTimestamp.Int64), obj.FillLevelVolume.Float64}, nil
}

func (r *autoWaterChangeResolver) BurnDown(ctx context.Context, obj *models.AutoWaterChange) ([]*model.FillLevel, error) {
	if !obj.FillLevelTimestamp.Valid || !obj.FillLevelVolume.Valid {
		return nil, nil
	}

	p := &models.Pump{ID: obj.FreshPumpID}
	doses, err := p.Doses(
		models.DoseWhere.Timestamp.GT(obj.FillLevelTimestamp.Int64),
	).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting pump doses: %w", err)
	}

	fillLevels := make([]*model.FillLevel, 0, len(doses))
	cumulative := float64(0)
	for _, dose := range doses {
		cumulative -= dose.Volume
		fillLevels = append(fillLevels, &model.FillLevel{
			Timestamp: int(dose.Timestamp),
			Volume:    obj.FillLevelVolume.Float64 - cumulative,
		})
	}

	return nil, nil
}

func (r *autoWaterChangeResolver) Events(ctx context.Context, obj *models.AutoWaterChange) ([]*models.AwcEvent, error) {
	events, err := obj.AwcEvents().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting AWC events: %w", err)
	}

	return events, nil
}

func (r *doseResolver) Message(ctx context.Context, obj *models.Dose) (*string, error) {
	if obj.Message.Valid {
		return &obj.Message.String, nil
	}
	return nil, nil
}

func (r *doserResolver) Name(ctx context.Context, obj *models.Doser) (*string, error) {
	if !obj.Name.Valid {
		return nil, nil
	}
	return &obj.Name.String, nil
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

func (r *firmataResolver) Name(ctx context.Context, obj *models.Firmata) (*string, error) {
	if !obj.Name.Valid {
		return nil, nil
	}
	return &obj.Name.String, nil
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

func (r *mutationResolver) CreateFirmata(ctx context.Context, serialPort string, baud int, name *string) (*models.Firmata, error) {
	m := &models.Firmata{
		ID:         uuid.New().String(),
		SerialPort: serialPort,
		Baud:       int64(baud),
		Name:       null.StringFromPtr(name),
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

func (r *mutationResolver) CreatePump(ctx context.Context, firmataID string, deviceID int, stepPin int, dirPin *int, enPin *int, acceleration *float64, name *string) (*models.Pump, error) {
	m := &models.Pump{
		ID:           uuid.New().String(),
		FirmataID:    firmataID,
		DeviceID:     int64(deviceID),
		StepPin:      int64(stepPin),
		Acceleration: null.Float64FromPtr(acceleration),
		Name:         null.StringFromPtr(name),
	}
	if dirPin != nil {
		m.DirPin = null.Int64From(int64(*dirPin))
	}
	if enPin != nil {
		m.EnPin = null.Int64From(int64(*enPin))
	}

	err := m.Insert(ctx, r.db, boil.Whitelist(
		models.PumpColumns.FirmataID,
		models.PumpColumns.DeviceID,
		models.PumpColumns.StepPin,
		models.PumpColumns.Acceleration,
		models.PumpColumns.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("inserting pump: %w", err)
	}

	r.firmatasController.Reset()
	r.atoController.Reset()
	r.awcController.Reset()

	return m, nil
}

func (r *mutationResolver) UpdatePump(ctx context.Context, id string, firmataID string, deviceID int, stepPin int, dirPin *int, enPin *int, acceleration *float64, name *string) (*models.Pump, error) {
	m := &models.Pump{
		ID:           id,
		FirmataID:    firmataID,
		DeviceID:     int64(deviceID),
		StepPin:      int64(stepPin),
		Acceleration: null.Float64FromPtr(acceleration),
		Name:         null.StringFromPtr(name),
	}
	if dirPin != nil {
		m.DirPin = null.Int64From(int64(*dirPin))
	}
	if enPin != nil {
		m.EnPin = null.Int64From(int64(*enPin))
	}

	_, err := m.Update(ctx, r.db, boil.Whitelist(
		models.PumpColumns.FirmataID,
		models.PumpColumns.DeviceID,
		models.PumpColumns.StepPin,
		models.PumpColumns.Acceleration,
		models.PumpColumns.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("inserting pump: %w", err)
	}

	r.firmatasController.Reset()
	r.atoController.Reset()
	r.awcController.Reset()

	return m, nil
}

func (r *mutationResolver) DeletePump(ctx context.Context, id string) (bool, error) {
	f := &models.Pump{ID: id}
	rows, err := f.Delete(ctx, r.db)
	if err != nil {
		return false, fmt.Errorf("deleting pump: %w", err)
	}

	r.firmatasController.Reset()
	r.atoController.Reset()
	r.awcController.Reset()

	return rows > 0, nil
}

func (r *mutationResolver) CalibratePump(ctx context.Context, pumpID string, steps int, volume float64) (*models.Calibration, error) {
	m := &models.Calibration{
		ID:        uuid.New().String(),
		PumpID:    pumpID,
		Timestamp: time.Now().Unix(),
		Steps:     int64(steps),
		Volume:    volume,
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

func (r *mutationResolver) CreateWaterLevelSensor(ctx context.Context, pin int, kind model.SensorKind, firmataID *string, detectionThreshold *int, invert bool, name *string) (*models.WaterLevelSensor, error) {
	m := &models.WaterLevelSensor{
		ID:        uuid.New().String(),
		Pin:       int64(pin),
		Kind:      kind.String(),
		FirmataID: null.StringFromPtr(firmataID),
		Invert:    invert,
		Name:      null.StringFromPtr(name),
	}
	if detectionThreshold != nil {
		m.DetectionThreshold = null.Int64From(int64(*detectionThreshold))
	}

	err := m.Insert(ctx, r.db, boil.Whitelist(
		models.WaterLevelSensorColumns.Pin,
		models.WaterLevelSensorColumns.Kind,
		models.WaterLevelSensorColumns.FirmataID,
		models.WaterLevelSensorColumns.Invert,
		models.WaterLevelSensorColumns.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("inserting water level sensor: %w", err)
	}

	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) UpdateWaterLevelSensor(ctx context.Context, id string, pin int, kind model.SensorKind, firmataID *string, detectionThreshold *int, invert bool, name *string) (*models.WaterLevelSensor, error) {
	m := &models.WaterLevelSensor{
		ID:        id,
		Pin:       int64(pin),
		Kind:      kind.String(),
		FirmataID: null.StringFromPtr(firmataID),
		Invert:    invert,
		Name:      null.StringFromPtr(name),
	}
	if detectionThreshold != nil {
		m.DetectionThreshold = null.Int64From(int64(*detectionThreshold))
	}

	_, err := m.Update(ctx, r.db, boil.Whitelist(
		models.WaterLevelSensorColumns.Pin,
		models.WaterLevelSensorColumns.Kind,
		models.WaterLevelSensorColumns.FirmataID,
		models.WaterLevelSensorColumns.Invert,
		models.WaterLevelSensorColumns.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("inserting water level sensor: %w", err)
	}

	r.firmatasController.Reset()
	r.atoController.Reset()
	r.awcController.Reset()

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

func (r *mutationResolver) CreateAutoTopOff(ctx context.Context, pumpID string, levelSensors []string, fillRate float64, fillInterval int, maxFillVolume float64, name *string) (*models.AutoTopOff, error) {
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

	m := &models.AutoTopOff{
		ID:            uuid.New().String(),
		PumpID:        pumpID,
		FillRate:      fillRate,
		FillInterval:  int64(fillInterval),
		MaxFillVolume: maxFillVolume,
		Name:          null.StringFromPtr(name),
	}
	waterLevelSensors := make([]*models.WaterLevelSensor, 0, len(levelSensors))
	for _, sensor := range levelSensors {
		waterLevelSensors = append(waterLevelSensors, &models.WaterLevelSensor{ID: sensor})
	}

	err = validateAutoTopOff(ctx, tx, m, waterLevelSensors)
	if err != nil {
		return nil, fmt.Errorf("validating auto top off: %w", err)
	}

	err = m.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting auto top off: %w", err)
	}

	err = m.SetWaterLevelSensors(ctx, tx, false, waterLevelSensors...)
	if err != nil {
		return nil, fmt.Errorf("associating water level sensor: %w", err)
	}

	r.atoController.Reset()
	r.firmatasController.Reset()

	return m, nil
}

func (r *mutationResolver) UpdateAutoTopOff(ctx context.Context, id string, pumpID string, levelSensors []string, fillRate float64, fillInterval int, maxFillVolume float64, name *string) (*models.AutoTopOff, error) {
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

	m := &models.AutoTopOff{
		ID:            id,
		PumpID:        pumpID,
		FillRate:      fillRate,
		FillInterval:  int64(fillInterval),
		MaxFillVolume: maxFillVolume,
		Name:          null.StringFromPtr(name),
	}
	waterLevelSensors := make([]*models.WaterLevelSensor, 0, len(levelSensors))
	for _, sensor := range levelSensors {
		waterLevelSensors = append(waterLevelSensors, &models.WaterLevelSensor{ID: sensor})
	}

	err = validateAutoTopOff(ctx, tx, m, waterLevelSensors)
	if err != nil {
		return nil, fmt.Errorf("validating auto top off: %w", err)
	}

	_, err = m.Update(ctx, tx, boil.Whitelist(
		models.AutoTopOffColumns.PumpID,
		models.AutoTopOffColumns.FillRate,
		models.AutoTopOffColumns.FillInterval,
		models.AutoTopOffColumns.MaxFillVolume,
		models.AutoTopOffColumns.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("inserting auto top off: %w", err)
	}

	err = m.SetWaterLevelSensors(ctx, tx, false, waterLevelSensors...)
	if err != nil {
		return nil, fmt.Errorf("associating water level sensor: %w", err)
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

func (r *mutationResolver) SetAutoTopOffEnabled(ctx context.Context, id string, enabled bool) (bool, error) {
	m := &models.AutoTopOff{
		ID:      id,
		Enabled: enabled,
	}

	_, err := m.Update(ctx, r.db, boil.Whitelist(models.AutoTopOffColumns.Enabled))
	if err != nil {
		return false, fmt.Errorf("updating auto top off: %w", err)
	}

	r.atoController.Reset()

	return enabled, nil
}

func (r *mutationResolver) SetATOFillLevel(ctx context.Context, id string, timestamp int, volume float64) (*models.AutoTopOff, error) {
	m, err := models.FindAutoTopOff(ctx, r.db, id)
	if err != nil {
		return nil, fmt.Errorf("getting ATO: %w", err)
	}

	m.FillLevelTimestamp = null.Int64From(int64(timestamp))
	m.FillLevelVolume = null.Float64From(volume)

	_, err = m.Update(ctx, r.db, boil.Whitelist(models.AutoTopOffColumns.FillLevelTimestamp, models.AutoTopOffColumns.FillLevelVolume))
	if err != nil {
		return nil, fmt.Errorf("updating auto top off: %w", err)
	}

	r.atoController.Reset()

	return m, nil
}

func (r *mutationResolver) CreateAutoWaterChange(ctx context.Context, freshPumpID string, wastePumpID string, exchangeRate float64, name *string) (*models.AutoWaterChange, error) {
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

	m := &models.AutoWaterChange{
		ID:           uuid.New().String(),
		FreshPumpID:  freshPumpID,
		WastePumpID:  wastePumpID,
		ExchangeRate: exchangeRate,
		Name:         null.StringFromPtr(name),
	}

	err = validateAutoWaterChange(ctx, tx, m)
	if err != nil {
		return nil, fmt.Errorf("validating auto top off: %w", err)
	}

	err = m.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("inserting auto water change: %w", err)
	}

	r.awcController.Reset()

	return m, nil
}

func (r *mutationResolver) UpdateAutoWaterChange(ctx context.Context, id string, freshPumpID string, wastePumpID string, exchangeRate float64, name *string) (*models.AutoWaterChange, error) {
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

	m := &models.AutoWaterChange{
		ID:           id,
		FreshPumpID:  freshPumpID,
		WastePumpID:  wastePumpID,
		ExchangeRate: exchangeRate,
		Name:         null.StringFromPtr(name),
	}

	err = validateAutoWaterChange(ctx, tx, m)
	if err != nil {
		return nil, fmt.Errorf("validating auto top off: %w", err)
	}

	_, err = m.Update(ctx, tx, boil.Whitelist(
		models.AutoWaterChangeColumns.FreshPumpID,
		models.AutoWaterChangeColumns.WastePumpID,
		models.AutoWaterChangeColumns.ExchangeRate,
		models.AutoWaterChangeColumns.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("updating auto water change: %w", err)
	}

	r.awcController.Reset()
	r.firmatasController.Reset()

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

func (r *mutationResolver) SetAutoWaterChangeEnabled(ctx context.Context, id string, enabled bool) (bool, error) {
	m := &models.AutoWaterChange{
		ID:      id,
		Enabled: enabled,
	}

	_, err := m.Update(ctx, r.db, boil.Whitelist(models.AutoTopOffColumns.Enabled))
	if err != nil {
		return false, fmt.Errorf("updating auto water change: %w", err)
	}

	r.awcController.Reset()

	return enabled, nil
}

func (r *mutationResolver) SetAWCFillLevel(ctx context.Context, id string, timestamp int, volume float64) (*models.AutoWaterChange, error) {
	m, err := models.FindAutoWaterChange(ctx, r.db, id)
	if err != nil {
		return nil, fmt.Errorf("getting AWC: %w", err)
	}

	m.FillLevelTimestamp = null.Int64From(int64(timestamp))
	m.FillLevelVolume = null.Float64From(volume)

	_, err = m.Update(ctx, r.db, boil.Whitelist(models.AutoWaterChangeColumns.FillLevelTimestamp, models.AutoWaterChangeColumns.FillLevelVolume))
	if err != nil {
		return nil, fmt.Errorf("updating auto water change: %w", err)
	}

	r.awcController.Reset()

	return m, nil
}

func (r *mutationResolver) CreateDoser(ctx context.Context, input model.DoserInput, name *string) (*models.Doser, error) {
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

func (r *mutationResolver) SetDoserEnabled(ctx context.Context, id string, enabled bool) (bool, error) {
	m := &models.Doser{
		ID:      id,
		Enabled: enabled,
	}

	_, err := m.Update(ctx, r.db, boil.Whitelist(models.AutoTopOffColumns.Enabled))
	if err != nil {
		return false, fmt.Errorf("updating doser: %w", err)
	}

	return enabled, nil
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

	if pump.Acceleration.Valid {
		err = firmata.StepperSetAcceleration(int(pump.DeviceID), float32(pump.Acceleration.Float64))
		if err != nil {
			return false, fmt.Errorf("setting pump acceleration: %w", err)
		}
	}

	err = firmata.StepperSetSpeed(int(pump.DeviceID), float32(math.Floor(speed)))
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

func (r *pumpResolver) Name(ctx context.Context, obj *models.Pump) (*string, error) {
	if !obj.Name.Valid {
		return nil, nil
	}
	return &obj.Name.String, nil
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
	calibration, err := obj.Calibrations(qm.OrderBy("timestamp DESC")).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting calibrations: %w", err)
	}

	return calibration, nil
}

func (r *pumpResolver) Acceleration(ctx context.Context, obj *models.Pump) (*float64, error) {
	if obj.Acceleration.Valid {
		return &obj.Acceleration.Float64, nil
	}
	return nil, nil
}

func (r *pumpResolver) History(ctx context.Context, obj *models.Pump) ([]*models.Dose, error) {
	doses, err := obj.Doses().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting dose history: %w", err)
	}

	return doses, nil
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

func (r *queryResolver) AutoTopOffs(ctx context.Context) ([]*models.AutoTopOff, error) {
	ms, err := models.AutoTopOffs().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting ATOs: %w", err)
	}

	return ms, nil
}

func (r *queryResolver) AutoTopOff(ctx context.Context, id string) (*models.AutoTopOff, error) {
	m, err := models.FindAutoTopOff(ctx, r.db, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting ATO: %w", err)
	}

	return m, nil
}

func (r *queryResolver) AutoWaterChanges(ctx context.Context) ([]*models.AutoWaterChange, error) {
	ms, err := models.AutoWaterChanges().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting AWCs: %w", err)
	}

	return ms, nil
}

func (r *queryResolver) AutoWaterChange(ctx context.Context, id string) (*models.AutoWaterChange, error) {
	m, err := models.FindAutoWaterChange(ctx, r.db, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting AWC: %w", err)
	}

	return m, nil
}

func (r *queryResolver) Dosers(ctx context.Context) ([]*models.Doser, error) {
	ms, err := models.Dosers().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("getting Dosers: %w", err)
	}

	return ms, nil
}

func (r *waterLevelSensorResolver) Name(ctx context.Context, obj *models.WaterLevelSensor) (*string, error) {
	if !obj.Name.Valid {
		return nil, nil
	}
	return &obj.Name.String, nil
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
	// Connect to the RPi
	rpi := raspi.NewAdaptor()
	err := rpi.Connect()
	if err != nil {
		return false, fmt.Errorf("Connecting to pi: %w", err)
	}

	return controller.WaterDetected(ctx, rpi, r.firmatasController, obj)
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

// Dose returns generated.DoseResolver implementation.
func (r *Resolver) Dose() generated.DoseResolver { return &doseResolver{r} }

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
type doseResolver struct{ *Resolver }
type doserResolver struct{ *Resolver }
type doserComponentResolver struct{ *Resolver }
type firmataResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type pumpResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type waterLevelSensorResolver struct{ *Resolver }
