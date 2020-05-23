package main

import (
	"log"
	"net/http"
)

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

// PostgresPlayerStore stores in the win in PostgreSQL
func (i *InMemoryPlayerStore) PostgresPlayerStore(name string, score int) int {
	return 0
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
