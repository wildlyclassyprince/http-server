package main

// NewInMemoryPlayerStore initializes store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore implements GetPlayerScore
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore returns the player's score given the name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin stores the win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
