package main

import (
	"fmt"

	"github.com/SaurabPoudel/smart-parking/types"
)

type Aggregator interface {
	AggregateInvoice(types.Invoice) error
}

type Storer interface {
	Insert(types.Invoice) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateInvoice(invoice types.Invoice) error {
	fmt.Printf("processing and inserting invoice for plate: %s, amount: %.2f\n", invoice.Plate, invoice.TotalAmount)
	return i.store.Insert(invoice)
}
