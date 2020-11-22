package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AnimeHandler handles the anime request (/anime/{id:[0-9]+})
func AnimeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Scraping anime list
	anime := Anime{}

	// If no cache available scrap data
	if len(cachedAnimeList) > 0 {
		for _, anm := range cachedAnimeList {
			if anm.ID == uint16(id) {
				anime = anm
			}
		}
	}

	// Setting up JSON headers
	w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode(anime)
}
