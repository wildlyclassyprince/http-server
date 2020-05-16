package main

import (
	"fmt"
	"net/http"
)

// PlayerServer takes the request sent in
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
