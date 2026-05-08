package main

import (
	"time"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next ParkingSessionServicer
}

func NewLogMiddleware(next ParkingSessionServicer) ParkingSessionServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) ProcessEntry(data types.ANPRData) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"plate":    data.Plate,
			"cameraID": data.CameraID,
			"event":    data.Event,
			"took":     time.Since(start),
			"err":      err,
		}).Info("Processing Entry")
	}(time.Now())

	err = m.next.ProcessEntry(data)
	return
}

func (m *LogMiddleware) ProcessExit(data types.ANPRData) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"plate":    data.Plate,
			"cameraID": data.CameraID,
			"event":    data.Event,
			"took":     time.Since(start),
			"err":      err,
		}).Info("Processing Exit")
	}(time.Now())

	err = m.next.ProcessExit(data)
	return
}

func (m *LogMiddleware) GetActiveSessions() map[string]*types.ParkingSession {
	return m.next.GetActiveSessions()
}

func (m *LogMiddleware) GetSessionFee(plate string) (fee float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"plate": plate,
			"fee":   fee,
			"took":  time.Since(start),
			"err":   err,
		}).Info("Getting Session Fee")
	}(time.Now())

	fee, err = m.next.GetSessionFee(plate)
	return
}
