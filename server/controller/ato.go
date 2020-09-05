package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/kerinin/doser/service/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ATO struct {
	eventCh  chan *models.AtoEvent
	db       *sql.DB
	firmatas *Firmatas
	reset    chan struct{}
}

func NewATO(db *sql.DB, firmatas *Firmatas) *ATO {
	return &ATO{
		eventCh:  make(chan *models.AtoEvent),
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

	wg.Add(1)
	go c.writeEvents(ctx, wg)

	jobs, err := c.setupJobs(ctx, wg)
	if err != nil {
		log.Printf("Failed to create initial ATO jobs: %s", err)
	}

	for {
		select {
		case <-c.reset:
			for _, job := range jobs {
				job.cancelFunc()
			}
			for _, job := range jobs {
				job.wg.Wait()
			}

			// NOTE: It's possible the new job is trying to talk to arduino at
			// the same time the old job is trying to shut down...
			jobs, err = c.setupJobs(ctx, wg)
			if err != nil {
				log.Printf("Failed to create ATO jobs: %s", err)
			}

		case <-ctx.Done():
			for _, job := range jobs {
				job.cancelFunc()
			}
			for _, job := range jobs {
				job.wg.Wait()
			}
			return
		}
	}
}

func (c *ATO) writeEvents(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case event := <-c.eventCh:
			if event == nil {
				return // channel closed
			}
			err := event.Insert(ctx, c.db, boil.Infer())
			if err != nil {
				log.Printf("Failed to persist ATO event %+v: %w", event, err)
			} else {
				log.Printf("ATO Event: %+v", event)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *ATO) setupJobs(ctx context.Context, wg *sync.WaitGroup) (jobs map[string]*Job, err error) {
	jobs = make(map[string]*Job)

	// If setup fails partway through, make sure we tear down any jobs that
	// were created before the failure
	defer func() {
		if err != nil {
			for _, job := range jobs {
				job.cancelFunc()
			}
			for _, job := range jobs {
				job.wg.Wait()
			}
		}
	}()

	atos, err := models.AutoTopOffs().All(ctx, c.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting ATOs: %w", err)
	}

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
		if err != nil {
			return nil, fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}

		var (
			jobCtx, cancel = context.WithCancel(ctx)
			jobWg          = &sync.WaitGroup{}
			job            = NewATOJob(c.eventCh, c.db, ato, pump, c.firmatas, firmata, sensors, calibration)
		)
		jobs[ato.ID] = &Job{cancelFunc: cancel, wg: jobWg}
		jobWg.Add(1)
		go job.Run(jobCtx, jobWg)
		log.Printf("Scheduled ATO job %s", ato.ID)
	}

	return jobs, nil
}
