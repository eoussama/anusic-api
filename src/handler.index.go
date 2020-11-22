package main

import (
	"encoding/json"
	"net/http"
)

// IndexHandler handles the entry request (/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Anusic API")
}
