package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eoussama/anusic-api/src/shared/utils"

	"github.com/gorilla/mux"
)

// AnimeHandler handles the anime request (/api/v1/anime/{id:[0-9]+})
func AnimeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Getting the respective anime
	anime := utils.Cache.GetAnimeByMALID(id)

	// Encoding the return value
	if anime != nil {
		json.NewEncoder(w).Encode(utils.FormatAnime(*anime))
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}
