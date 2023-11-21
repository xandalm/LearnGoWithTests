package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		sync.RWMutex{},
		map[string]int{},
	}
}

type InMemoryPlayerStore struct {
	mu    sync.RWMutex
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.store[name]
}
