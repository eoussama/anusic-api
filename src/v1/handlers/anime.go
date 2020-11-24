package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eoussama/anusic-api/src/v1/models"
	"github.com/eoussama/anusic-api/src/v1/utils"

	"github.com/gorilla/mux"
)

// AnimeHandler handles the anime request (/anime/{id:[0-9]+})
func AnimeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Scraping anime list
	anime := models.Anime{}

	// If no cache available scrap data
	if len(utils.CachedAnimeList) > 0 {
		for _, anm := range utils.CachedAnimeList {
			if anm.ID == uint16(id) {
				anime = anm
			}
		}
	}

	// Setting up JSON headers
	// w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode(anime)
}
