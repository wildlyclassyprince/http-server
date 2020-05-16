package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer implements the `Handler` by type casting. It takes an HTTP request
// to GET a player's score given the name, and POSTS (prints) the player's score.
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
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
