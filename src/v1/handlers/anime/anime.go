package anime

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
		w.WriteHeader(http.StatusNotFound)
		utils.ReturnResponse(w, r, nil, models.Error.AnimeNotFound(models.Error{}, id))
	} else {
		utils.ReturnResponse(w, r, utils.FormatAnime(*anime), nil)
	}
}
