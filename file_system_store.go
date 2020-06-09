package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore stores database objects
type FileSystemPlayerStore struct {
	database io.Reader
}

// GetLeague stores player scores in JSON file
func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.database).Decode(&league)
	return league
}
