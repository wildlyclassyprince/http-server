package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores the retrieved player's score
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer implements the handler method 'ServeHTTP' for a 'PlayerStore' interface.
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP is the handler that processes the HTTP requests
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// GetPlayerScore takes the player name and returns the score
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
