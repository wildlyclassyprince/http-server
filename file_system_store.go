package main

import (
	"io"
)

// FileSystemPlayerStore stores database objects
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague stores player scores in JSON file
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore returns the player's score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}
