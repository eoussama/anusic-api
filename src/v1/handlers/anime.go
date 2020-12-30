package handlers

import (
	"net/http"
	"strconv"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"

	"github.com/gorilla/mux"
)

// AnimeHandler handles the anime request (/api/v1/anime/{id:[0-9]+})
func AnimeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Getting the respective anime
	anime := utils.Cache.GetAnimeByMALID(id)

	// Returning the response value
	if anime == nil {
		utils.ReturnResponse(w, nil, models.Error.AnimeNotFound(models.Error{}, id))
	} else {
		utils.ReturnResponse(w, utils.FormatAnime(*anime), nil)
	}
}
