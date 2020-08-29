package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/kerinin/doser/service/models"
)

type AWC struct {
	eventCh  chan<- Event
	db       *sql.DB
	firmatas *Firmatas
	reset    chan struct{}
	jobs     map[string]context.CancelFunc
}

func NewAWC(eventCh chan<- Event, db *sql.DB, firmatas *Firmatas) *AWC {
	return &AWC{
		eventCh:  eventCh,
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

	jobs, err := c.setupJobs(ctx, wg)
	if err != nil {
		log.Printf("Failed to create initial AWC jobs: %s", err)
	}

	for {
		select {
		case <-c.reset:
			for _, cancel := range jobs {
				cancel()
			}

			// NOTE: It's possible the new job is trying to talk to arduino at
			// the same time the old job is trying to shut down...
			jobs, err = c.setupJobs(ctx, wg)
			if err != nil {
				log.Printf("Failed to create AWC jobs: %s", err)
			}

		case <-ctx.Done():
			for _, cancel := range jobs {
				cancel()
			}
			return
		}
	}
}

func (c *AWC) setupJobs(ctx context.Context, wg *sync.WaitGroup) (jobs map[string]context.CancelFunc, err error) {
	jobs = make(map[string]context.CancelFunc)

	// If setup fails partway through, make sure we tear down any jobs that
	// were created before the failure
	defer func() {
		if err != nil {
			for _, cancel := range jobs {
				cancel()
			}
		}
	}()

	awcs, err := models.AutoWaterChanges().All(ctx, c.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("getting AWCs: %w", err)
	}

	for _, awc := range awcs {
		// Fetch resources necessary for the ATO job
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
			return nil, fmt.Errorf("configuring fresh pump: %w", err)
		}
		wasteFirmata, err := c.firmatas.Get(ctx, wastePump.FirmataID)
		if err != nil {
			return nil, fmt.Errorf("getting fresh pump firmata: %w", err)
		}
		err = ConfigurePump(wastePump, wasteFirmata)
		if err != nil {
			return nil, fmt.Errorf("configuring waste pump: %w", err)
		}
		freshCalibration, err := freshPump.Calibrations().One(ctx, c.db)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("refusing to run ATO job with uncalibrated fresh pump")
		} else if err != nil {
			return nil, fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}
		wasteCalibration, err := wastePump.Calibrations().One(ctx, c.db)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("refusing to run ATO job with uncalibrated waste pump")
		} else if err != nil {
			return nil, fmt.Errorf("getting pump calibration (aborting job run): %w", err)
		}
		var (
			jobCtx, cancel = context.WithCancel(ctx)
			job            = NewAWCJob(c.eventCh, awc, freshPump, wastePump, freshFirmata, wasteFirmata, freshCalibration, wasteCalibration)
		)
		jobs[awc.ID] = cancel
		wg.Add(1)
		go job.Run(jobCtx, wg)
	}

	return jobs, nil
}
