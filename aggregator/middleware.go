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

func (m *LogMiddleware) AggregateInvoice(invoice types.Invoice) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"plate":  invoice.Plate,
			"amount": invoice.TotalAmount,
			"took":   time.Since(start),
			"err":    err,
		}).Info("Aggregating invoice")
	}(time.Now())

	err = m.next.AggregateInvoice(invoice)
	return err
}
