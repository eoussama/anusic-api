package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eoussama/anusic-api/config"
	"github.com/eoussama/anusic-api/src/shared/models"
)

// IndexHandler handles the entry request (/api/v1/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			models.Response
			Data struct {
				config.Meta
				Version uint8 `json:"version"`
			} `json:"data"`
		}{
			models.Response{},
			struct {
				config.Meta
				Version uint8 `json:"version"`
			}{
				config.Info,
				uint8(1),
			},
		},
	)
}
