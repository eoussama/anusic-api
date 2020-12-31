package handlers

import (
	"net/http"

	"github.com/eoussama/anusic-api/config"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// IndexHandler handles the entry request (/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	utils.ReturnResponse(w, r, config.Info, nil, http.StatusOK)
}
