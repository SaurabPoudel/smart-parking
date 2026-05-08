package main

import (
	"time"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateParking(p types.Parking) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("Aggregating parking data")
	}(time.Now())

	err = m.next.AggregateParking(p)
	return err
}
