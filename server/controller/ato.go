package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/kerinin/doser/service/models"
	"github.com/robfig/cron/v3"
)

type ATO struct {
	eventCh  chan<- Event
	db       *sql.DB
	firmatas *Firmatas
	reset    chan struct{}
}

func NewATO(eventCh chan<- Event, db *sql.DB, firmatas *Firmatas) *ATO {
	return &ATO{
		eventCh:  eventCh,
		db:       db,
		firmatas: firmatas,
		reset:    make(chan struct{}, 1),
	}
}

func (c *ATO) Reset() {
	c.reset <- struct{}{}
}

func (c *ATO) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	crn, err := c.setupCron(ctx, wg)
	if err != nil {
		log.Printf("Failed to create initial ATO jobs: %s", err)
	}

	crn.Start()
	defer crn.Stop()

	for {
		select {
		case <-c.reset:
			nextCrn, err := c.setupCron(ctx, wg)
			if err != nil {
				log.Printf("Failed to refresh ATO jobs: %s", err)
				continue
			}

			// Don't stop the running cron unless the new cron was created successfully
			if crn != nil {
				stopCtx := crn.Stop()
				// NOTE: Do we want to wait for running jobs to terminate?
				<-stopCtx.Done()
			}

			// Start processing the next cron
			if nextCrn != nil {
				nextCrn.Start()
			}
			crn = nextCrn

		case <-ctx.Done():
			return
		}
	}
}

func (c *ATO) setupCron(ctx context.Context, wg *sync.WaitGroup) (*cron.Cron, error) {
	atos, err := models.AutoTopOffs().All(ctx, c.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting ATOs: %w", err)
	}

	crn := cron.New()
	for _, ato := range atos {
		// Fetch resources necessary for the ATO job
		pump, err := ato.Pump().One(ctx, c.db)
		if err != nil {
			return nil, fmt.Errorf("getting pump (aborting job run): %w", err)
		}
		firmata, err := c.firmatas.Get(ctx, pump.FirmataID)
		if err != nil {
			return nil, fmt.Errorf("getting pump firmata: %w", err)
		}
		err = ConfigurePump(pump, firmata)
		if err != nil {
			return nil, fmt.Errorf("configuring pump: %w", err)
		}
		sensors, err := ato.WaterLevelSensors().All(ctx, c.db)
		if err != nil {
			return nil, fmt.Errorf("getting sensors (aborting job run): %w", err)
		}
		calibration, err := pump.Calibrations().One(ctx, c.db)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("refusing to run ATO job with uncalibrated pump")
		} else if err != nil {
			return nil, fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}

		_, err = crn.AddJob(ato.FillFrequency, NewATOJob(ctx, wg, c.eventCh, ato, pump, firmata, sensors, calibration))
		if err != nil {
			return nil, fmt.Errorf("adding cron job: %w", err)
		}
		log.Printf("Scheduled ATO job %s", ato.ID)
	}

	return crn, nil
}
