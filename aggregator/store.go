package main

import (
	"fmt"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/sirupsen/logrus"
)

type MemoryStore struct {
	invoices map[string]*types.Invoice
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		invoices: make(map[string]*types.Invoice),
	}
}

func (m *MemoryStore) Insert(invoice types.Invoice) error {
	if invoice.Plate == "" {
		return fmt.Errorf("plate cannot be empty")
	}

	if existing, exists := m.invoices[invoice.Plate]; exists {
		existing.TotalSessions += invoice.TotalSessions
		existing.TotalAmount += invoice.TotalAmount
		existing.PeriodEnd = invoice.PeriodEnd
		logrus.Infof("Updated invoice in store for plate: %s", invoice.Plate)
	} else {
		m.invoices[invoice.Plate] = &invoice
		logrus.Infof("Inserted new invoice in store for plate: %s", invoice.Plate)
	}

	return nil
}

func (m *MemoryStore) GetAll() []types.Invoice {
	invoices := make([]types.Invoice, 0)
	for _, invoice := range m.invoices {
		invoices = append(invoices, *invoice)
	}
	return invoices
}
