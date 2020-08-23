package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
	"github.com/robfig/cron/v3"
)

type ATO struct {
	eventCh  chan<- Event
	db       *sql.DB
	firmatas map[string]*gomata.Firmata
	reset    chan struct{}
}

func NewATO(eventCh chan<- Event, db *sql.DB, firmatas map[string]*gomata.Firmata) *ATO {
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
		_, err := crn.AddJob(ato.FillFrequency, NewATOJob(ctx, wg, c.eventCh, c.db, c.firmatas, ato))
		if err != nil {
			return nil, fmt.Errorf("adding cron job: %w", err)
		}
	}

	return crn, nil
}
