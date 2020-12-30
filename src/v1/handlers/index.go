package handlers

import (
	"net/http"

	"github.com/eoussama/anusic-api/config"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// IndexHandler handles the entry request (/api/v1/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	utils.ReturnResponse(w, struct {
		config.Meta
		Version uint8
	}{
		config.Info,
		uint8(1),
	})
}
