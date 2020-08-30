package controller

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huin/goserial"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

type Firmatas struct {
	eventCh  chan<- Event
	db       *sql.DB
	firmatas map[string]*gomata.Firmata
	reset    chan struct{}
}

func NewFirmatas(eventCh chan<- Event, db *sql.DB) *Firmatas {
	return &Firmatas{eventCh, db, make(map[string]*gomata.Firmata), make(chan struct{}, 1)}
}

func (c *Firmatas) Reset() error {
	for _, f := range c.firmatas {
		err := f.Disconnect
		if err != nil {
			return fmt.Errorf("disconnecting firmata: %w", err)
		}
	}

	c.firmatas = make(map[string]*gomata.Firmata)
	return nil
}

func (c *Firmatas) Get(ctx context.Context, firmataID string) (*gomata.Firmata, error) {
	if f, found := c.firmatas[firmataID]; found {
		return f, nil
	}

	firmata, err := models.FindFirmata(ctx, c.db, firmataID)
	if err != nil {
		return nil, fmt.Errorf("getting firmata from DB: %w", err)
	}

	sensors, err := firmata.FirmatumWaterLevelSensors().All(ctx, c.db)
	if err != nil {
		return nil, fmt.Errorf("getting sensors from DB: %w", err)
	}

	config := &goserial.Config{Name: firmata.SerialPort, Baud: int(firmata.Baud)}
	port, err := goserial.OpenPort(config)
	if err != nil {
		return nil, fmt.Errorf("opening serial port: %w", err)
	}

	f := gomata.New()
	f.Connect(port)
	for _, sensor := range sensors {
		err = f.ReportDigital(int(sensor.Pin), 1)
		if err != nil {
			return nil, fmt.Errorf("requesting digital reports for pin %d: %w", sensor.Pin, err)
		}
	}
	c.firmatas[firmataID] = f

	return f, nil
}
