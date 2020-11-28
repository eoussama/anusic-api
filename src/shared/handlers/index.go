package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eoussama/anusic-api/config"
)

// IndexHandler handles the entry request (/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(config.Info)
}
