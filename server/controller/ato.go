package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kerinin/doser/service/models"
	"github.com/robfig/cron/v3"
)

type ATOControl struct {
	db    *sql.DB
	cron  *cron.Cron
	reset chan struct{}
}

func NewATOControl(db *sql.DB) *ATOControl {
	return &ATOControl{db: db, reset: make(chan struct{}, 0)}
}

func (c *ATOControl) Reset() {
	c.reset <- struct{}{}
}

func (c *ATOControl) Run(ctx context.Context) {
	crn, err := c.setupCron(ctx)
	if err != nil {
		log.Printf("Failed to create initial ATO jobs: %s", err)
	}

	crn.Start()
	defer c.cron.Stop()

	for {
		select {
		case <-c.reset:
			nextCrn, err := c.setupCron(ctx)
			if err != nil {
				log.Printf("Failed to refresh ATO jobs: %s", err)
				continue
			}
			stopCtx := crn.Stop()
			// NOTE: Do we want to wait for running jobs to terminate?
			<-stopCtx.Done()
			nextCrn.Start()
			crn = nextCrn

		case <-ctx.Done():
			return
		}
	}
}

func (c *ATOControl) setupCron(ctx context.Context) (*cron.Cron, error) {
	atos, err := models.AutoTopOffs().All(ctx, c.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting ATOs: %w", err)
	}

	crn := cron.New()
	for _, ato := range atos {
		_, err := crn.AddJob(ato.FillFrequency, NewATOJob(ato))
		if err != nil {
			return nil, fmt.Errorf("adding cron job: %w", err)
		}
	}

	return crn, nil
}
