package main

import "sync"

// NewInMemoryPlayerStore initialize an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex // A mutex is used to synchronize read/write access to the map
}

// GetPlayerScore retrieves scores for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

// RecordWin will record a player's win.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

// GetLeague currently doesn't work, but it should return the player league.
func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
