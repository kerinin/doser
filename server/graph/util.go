package graph

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kerinin/doser/service/models"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NullInt64ToIntPtr(v null.Int64) *int {
	if !v.Valid {
		return nil
	}
	vNew := int(v.Int64)
	return &vNew
}

func validateAutoTopOff(ctx context.Context, tx *sql.Tx, ato *models.AutoTopOff, sensors []*models.WaterLevelSensor) error {
	intervalVolume := float64(ato.FillInterval) * ato.FillRate

	if intervalVolume < ato.MaxFillVolume {
		return fmt.Errorf("maximum possible job fill volume %d is less than requested maximum fill volume %d", intervalVolume, ato.MaxFillVolume)
	}

	pump, err := models.FindPump(ctx, tx, ato.PumpID)
	if err != nil {
		return fmt.Errorf("finding ATO pump: %w", err)
	}

	_, err = pump.Calibrations(qm.OrderBy("timestamp DESC")).One(ctx, tx)
	if err == sql.ErrNoRows {
		return fmt.Errorf("refusing to create ATO with uncalibrated pump")
	} else if err != nil {
		return fmt.Errorf("getting pump calibration: %w", err)
	}

	if len(sensors) == 0 {
		return fmt.Errorf("missing water level sensors")
	}

	return nil
}

func validateAutoWaterChange(ctx context.Context, tx *sql.Tx, awc *models.AutoWaterChange) error {
	if awc.ExchangeRate < 0 {
		return fmt.Errorf("Exchange rate must be positive")
	}

	if (awc.ExchangeRate*1000)-awc.SalinityAdjustment < 0 {
		return fmt.Errorf("Salinity adjustment would produce negative waste water pump rate")
	}

	freshPump, err := models.FindPump(ctx, tx, awc.FreshPumpID)
	if err != nil {
		return fmt.Errorf("finding fresh pump: %w", err)
	}

	_, err = freshPump.Calibrations(qm.OrderBy("timestamp DESC")).One(ctx, tx)
	if err == sql.ErrNoRows {
		return fmt.Errorf("refusing to create AWC with uncalibrated fresh pump")
	} else if err != nil {
		return fmt.Errorf("getting pump calibration: %w", err)
	}

	wastePump, err := models.FindPump(ctx, tx, awc.WastePumpID)
	if err != nil {
		return fmt.Errorf("finding waste pump: %w", err)
	}

	_, err = wastePump.Calibrations(qm.OrderBy("timestamp DESC")).One(ctx, tx)
	if err == sql.ErrNoRows {
		return fmt.Errorf("refusing to create AWC with uncalibrated waste pump")
	} else if err != nil {
		return fmt.Errorf("getting pump calibration: %w", err)
	}

	return nil
}
