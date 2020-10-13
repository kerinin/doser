package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

type Firmatas struct {
	db       *sql.DB
	firmatas map[string]*gomata.Firmata
	reset    chan struct{}
	mx       *sync.Mutex
}

func NewFirmatas(db *sql.DB) *Firmatas {
	return &Firmatas{db, make(map[string]*gomata.Firmata), make(chan struct{}, 1), &sync.Mutex{}}
}

func (c *Firmatas) Reset() error {
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, f := range c.firmatas {
		err := f.Disconnect()
		if err != nil {
			log.Printf("Failed to disconnect firmata: %w", err)
		}
	}

	c.firmatas = make(map[string]*gomata.Firmata)
	return nil
}

func (c *Firmatas) Get(ctx context.Context, firmataID string) (*gomata.Firmata, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

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

	config := serial.OpenOptions{
		PortName:              firmata.SerialPort,
		BaudRate:              uint(firmata.Baud),
		DataBits:              8,
		StopBits:              1,
		InterCharacterTimeout: 60 * 1000,
	}
	port, err := serial.Open(config)
	if err != nil {
		return nil, fmt.Errorf("opening serial port: %w", err)
	}

	log.Printf("Connecting to firmata at port %s", firmata.SerialPort)

	f := gomata.New()
	connected := make(chan error)
	go func() {
		connected <- f.Connect(port)
	}()

	select {
	case err := <-connected:
		if err != nil {
			return nil, fmt.Errorf("Connecting to firmata: %w", err)
		}

		for _, sensor := range sensors {
			if sensor.DetectionThreshold.Valid {
				log.Printf("Requesting analog reports for firmata %s pin %d", firmataID, sensor.Pin)

				err := f.SetPinMode(int(sensor.Pin), gomata.AnalogPin)
				if err != nil {
					return nil, fmt.Errorf("setting pin mode: %w", err)
				}

				err = f.ReportAnalog(int(sensor.Pin), 1)
				if err != nil {
					return nil, fmt.Errorf("requesting analog reports for pin %d: %w", sensor.Pin, err)
				}
			} else {

				log.Printf("Requesting digital reports for firmata %s pin %d", firmataID, sensor.Pin)

				err := f.SetPinMode(int(sensor.Pin), gomata.PullupPin)
				if err != nil {
					return nil, fmt.Errorf("setting pin mode: %w", err)
				}

				err = f.ReportDigital(int(sensor.Pin), 1)
				if err != nil {
					return nil, fmt.Errorf("requesting digital reports for pin %d: %w", sensor.Pin, err)
				}
			}
		}
		c.firmatas[firmataID] = f

		return f, nil

	case <-time.After(30 * time.Second):
		err := f.Disconnect()
		if err != nil {
			log.Printf("Failed to disconnect from firmata after failing to connect: %s", err)
		}
		return nil, fmt.Errorf("Failed to connect to firmata")
	}
}
