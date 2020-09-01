package graph

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kerinin/doser/service/models"
	null "github.com/volatiletech/null/v8"
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

	_, err = pump.Calibrations().One(ctx, tx)
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
