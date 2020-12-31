package handlers

import (
	"net/http"

	"github.com/eoussama/anusic-api/config"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// IndexHandler handles the entry request (/api/v1/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	info := struct {
		config.Meta
		Version uint8 `json:"version"`
	}{config.Info, uint8(1)}

	// Returning the response value
	utils.ReturnResponse(w, r, info, nil)
}
