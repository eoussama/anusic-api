package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eoussama/anusic-api/config"
)

// IndexHandler handles the entry request (/api/v1/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		config.Meta
		Version uint8 `json:"version"`
	}{
		config.Info,
		uint8(1),
	})
}
