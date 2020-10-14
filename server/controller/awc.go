package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/kerinin/doser/service/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AWC struct {
	eventCh  chan *models.AwcEvent
	db       *sql.DB
	firmatas *Firmatas
	reset    chan struct{}
}

func NewAWC(db *sql.DB, firmatas *Firmatas) *AWC {
	return &AWC{
		eventCh:  make(chan *models.AwcEvent),
		db:       db,
		firmatas: firmatas,
		reset:    make(chan struct{}, 1),
	}
}

func (c *AWC) Reset() {
	c.reset <- struct{}{}
}

func (c *AWC) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go c.writeEvents(ctx, wg)

	jobs, err := c.setupJobs(ctx, wg)
	if err != nil {
		log.Printf("Failed to create initial AWC jobs: %s", err)
	}

	for {
		select {
		case <-c.reset:
			for _, job := range jobs {
				job.cancelFunc()
			}

			// NOTE: It's possible the new job is trying to talk to arduino at
			// the same time the old job is trying to shut down...
			jobs, err = c.setupJobs(ctx, wg)
			if err != nil {
				log.Printf("Failed to create AWC jobs: %s", err)
				c.Reset()
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

func (c *AWC) writeEvents(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case event := <-c.eventCh:
			if event == nil {
				return // channel closed
			}
			err := event.Insert(ctx, c.db, boil.Infer())
			if err != nil {
				log.Printf("Failed to persist AWC event %+v: %w", event, err)
			} else {
				log.Printf("AWC Event: %+v", event)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *AWC) setupJobs(ctx context.Context, wg *sync.WaitGroup) (jobs map[string]*Job, err error) {
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

	awcs, err := models.AutoWaterChanges(models.AutoWaterChangeWhere.Enabled.EQ(true)).All(ctx, c.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting AWCs: %w", err)
	}

	for _, awc := range awcs {
		// Fetch resources necessary for the AWC job
		freshPump, err := awc.FreshPump().One(ctx, c.db)
		if err != nil {
			return nil, fmt.Errorf("getting fresh water pump (aborting job run): %w", err)
		}
		wastePump, err := awc.WastePump().One(ctx, c.db)
		if err != nil {
			return nil, fmt.Errorf("getting waste water pump (aborting job run): %w", err)
		}
		freshFirmata, err := c.firmatas.Get(ctx, freshPump.FirmataID)
		if err != nil {
			return nil, fmt.Errorf("getting fresh pump firmata: %w", err)
		}
		err = ConfigurePump(freshPump, freshFirmata)
		if err != nil {
			c.firmatas.Reset()
			return nil, fmt.Errorf("configuring fresh pump: %w", err)
		}
		wasteFirmata, err := c.firmatas.Get(ctx, wastePump.FirmataID)
		if err != nil {
			return nil, fmt.Errorf("getting fresh pump firmata: %w", err)
		}
		err = ConfigurePump(wastePump, wasteFirmata)
		if err != nil {
			c.firmatas.Reset()
			return nil, fmt.Errorf("configuring waste pump: %w", err)
		}
		freshCalibration, err := freshPump.Calibrations(qm.OrderBy(models.CalibrationColumns.Timestamp)).One(ctx, c.db)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("refusing to run AWC job with uncalibrated fresh pump")
		} else if err != nil {
			return nil, fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}
		wasteCalibration, err := wastePump.Calibrations(qm.OrderBy(models.CalibrationColumns.Timestamp)).One(ctx, c.db)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("refusing to run AWC job with uncalibrated waste pump")
		} else if err != nil {
			return nil, fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}

		var (
			jobCtx, cancel = context.WithCancel(ctx)
			jobWg          = &sync.WaitGroup{}
			job            = NewAWCJob(c, awc, freshPump, wastePump, freshFirmata, wasteFirmata, freshCalibration, wasteCalibration)
		)
		jobs[awc.ID] = &Job{cancel, jobWg}
		jobWg.Add(1)
		go job.Run(jobCtx, jobWg)
		log.Printf("Scheduled AWC job %s", awc.ID)
	}

	return jobs, nil
}
