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
		var (
			jobCtx, cancel = context.WithCancel(ctx)
			job            = NewAWCJob(c.db, c.firmatas, awc)
		)
		c.jobs[awc.ID] = cancel
		wg.Add(1)
		go job.Run(jobCtx, wg)
	}

	return nil
}
