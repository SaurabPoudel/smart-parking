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
		logFields := logrus.Fields{
			"plate":  invoice.Plate,
			"amount": invoice.TotalAmount,
			"took":   time.Since(start),
		}

		if err != nil {
			logFields["err"] = err
			logrus.WithFields(logFields).Error("Aggregating invoice failed")
		} else {
			logrus.WithFields(logFields).Info("Aggregating invoice")
		}
	}(time.Now())

	return m.next.AggregateInvoice(invoice)
}
