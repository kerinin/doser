package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

type AWCController struct {
	db       *sql.DB
	firmatas map[string]*gomata.Firmata
	reset    chan struct{}
	jobs     map[string]context.CancelFunc
}

func NewAWCController() *AWCController {
	return &AWCController{}
}

func (c *AWCController) Run(ctx context.Context, wg *sync.WaitGroup) {
	// TODO: subsequent runs...
	err := c.setupJobs(ctx, wg)
	if err != nil {
		log.Printf("Failed to create initial AWC jobs: %s", err)
	}

	for {
		select {
		case <-c.reset:
		case <-ctx.Done():
			return
		}
	}
}

func (c *AWCController) setupJobs(ctx context.Context, wg *sync.WaitGroup) error {
	awcs, err := models.AutoWaterChanges().All(ctx, c.db)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return fmt.Errorf("getting AWCs: %w", err)
	}

	for k, cancel := range c.jobs {
		cancel()
		delete(c.jobs, k)
	}

	for _, awc := range awcs {
		// Fetch resources necessary for the ATO job
		freshPump, err := awc.FreshPump().One(ctx, c.db)
		if err != nil {
			return fmt.Errorf("getting fresh water pump (aborting job run): %w", err)
		}
		wastePump, err := awc.WastePump().One(ctx, c.db)
		if err != nil {
			return fmt.Errorf("getting waste water pump (aborting job run): %w", err)
		}
		freshFirmata, found := c.firmatas[freshPump.FirmataID]
		if !found {
			return fmt.Errorf("unrecognized firmata ID %s for fresh pump %s (aborting job run)", freshPump.FirmataID, freshPump.ID)
		}
		wasteFirmata, found := c.firmatas[wastePump.FirmataID]
		if !found {
			return fmt.Errorf("unrecognized firmata ID %s for waste pump %s (aborting job run)", wastePump.FirmataID, wastePump.ID)
		}
		freshCalibration, err := freshPump.Calibrations().One(ctx, c.db)
		if err == sql.ErrNoRows {
			return fmt.Errorf("refusing to run ATO job with uncalibrated fresh pump")
		} else if err != nil {
			return fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}
		wasteCalibration, err := wastePump.Calibrations().One(ctx, c.db)
		if err == sql.ErrNoRows {
			return fmt.Errorf("refusing to run ATO job with uncalibrated waste pump")
		} else if err != nil {
			return fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}
		var (
			jobCtx, cancel = context.WithCancel(ctx)
			job            = NewAWCJob(awc, freshPump, wastePump, freshFirmata, wasteFirmata, freshCalibration, wasteCalibration)
		)
		c.jobs[awc.ID] = cancel
		wg.Add(1)
		go job.Run(jobCtx, wg)
	}

	return nil
}
