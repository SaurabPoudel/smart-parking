package main

import "github.com/SaurabPoudel/smart-parking/types"

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Insert(p types.Parking) error {
	return nil
}
