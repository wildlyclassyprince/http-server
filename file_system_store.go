package main

import (
	"io"
)

// FileSystemPlayerStore stores database objects
type FileSystemPlayerStore struct {
	database io.Reader
}

// GetLeague stores player scores in JSON file
func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
