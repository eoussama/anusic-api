package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eoussama/anusic-api/config"
	"github.com/eoussama/anusic-api/src/shared/models"
)

// IndexHandler handles the entry request (/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			models.Response
			Data config.Meta `json:"data"`
		}{
			models.Response{},
			config.Info,
		},
	)
}
