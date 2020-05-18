package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore implements GetPlayerScore
type InMemoryPlayerStore struct{}

// GetPlayerScore returns the player's score given the name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin stores the win
func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
