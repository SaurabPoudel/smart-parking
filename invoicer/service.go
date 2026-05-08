package main

import (
	"fmt"
	"time"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Invoicer interface {
	AggregateInvoice(session types.ParkingSession) error
	GetInvoiceByPlate(plate string) (types.Invoice, error)
	GetAllInvoices() []types.Invoice
}

type InvoiceService struct {
	invoices map[string]*types.Invoice
}

func NewInvoiceService() Invoicer {
	return &InvoiceService{
		invoices: make(map[string]*types.Invoice),
	}
}

func (is *InvoiceService) AggregateInvoice(session types.ParkingSession) error {
	if session.Plate == "" {
		return fmt.Errorf("plate cannot be empty")
	}

	if existing, exists := is.invoices[session.Plate]; exists {
		// Aggregate with existing invoice
		existing.TotalSessions++
		existing.TotalAmount += session.Fee
		existing.PeriodEnd = time.Now()
		logrus.Infof("Updated invoice for plate: %s, total amount: %.2f", session.Plate, existing.TotalAmount)
	} else {
		// Create new invoice
		invoice := &types.Invoice{
			InvoiceID:     uuid.New().String(),
			Plate:         session.Plate,
			TotalSessions: 1,
			TotalAmount:   session.Fee,
			PeriodStart:   time.Now(),
			PeriodEnd:     time.Now(),
			Status:        "PENDING",
		}
		is.invoices[session.Plate] = invoice
		logrus.Infof("Created new invoice for plate: %s, amount: %.2f", session.Plate, session.Fee)
	}

	return nil
}

func (is *InvoiceService) GetInvoiceByPlate(plate string) (types.Invoice, error) {
	invoice, exists := is.invoices[plate]
	if !exists {
		return types.Invoice{}, fmt.Errorf("invoice not found for plate: %s", plate)
	}
	return *invoice, nil
}

func (is *InvoiceService) GetAllInvoices() []types.Invoice {
	invoices := make([]types.Invoice, 0)
	for _, invoice := range is.invoices {
		invoices = append(invoices, *invoice)
	}
	return invoices
}
