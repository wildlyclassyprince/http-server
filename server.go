package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

const (
	// Host is the hostname
	Host = "localhost"
	// Port is the default PostgreSQL port
	Port = 5432
	// User is the default name "postgres"
	User = "postgres"
	// Password is the "postgres" password
	Password = ""
	// Dbname is the name of the database
	Dbname = "postgres"
)

// PlayerStore stores the retrieved player's score
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	PostgresPlayerStore(name string, score int) int
}

// PlayerServer implements the handler method 'ServeHTTP' for a 'PlayerStore' interface.
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP is the handler that processes the HTTP requests
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
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

// PostgresPlayerStore establishes a database connection
// and stores player scores in a PostgreSQL database
func PostgresPlayerStore(name string, score int) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
INSERT INTO players (name, score)
VALUES ($1, $2)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, name, score).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
