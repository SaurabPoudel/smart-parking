package main

import (
	"fmt"

	"github.com/SaurabPoudel/smart-parking/types"
)

type Aggregator interface {
	AggregateParking(types.Parking) error
}

type Storer interface {
	Insert(types.Parking) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateParking(p types.Parking) error {
	fmt.Println("processing and inserting parking data for ", p)
	return i.store.Insert(p)
}
