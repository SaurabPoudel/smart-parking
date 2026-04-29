package main

import (
	"time"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.ANPRData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"cameraID": data.CameraID,
			"plate":    data.Plate,
			"Event":    data.Event,
			"took":     time.Since(start),
		}).Info("producing to Kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
