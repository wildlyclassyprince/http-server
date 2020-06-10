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
