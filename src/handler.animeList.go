package main

import (
	"encoding/json"
	"net/http"
)

// AnimeListHandler handles the anime list request (/)
func AnimeListHandler(w http.ResponseWriter, r *http.Request) {

	// Scraping anime list
	animeTitles := cachedAnimeList

	// If no cache available scrap data
	if len(animeTitles) == 0 {
		animeTitles = scrapAnimeList()
	}

	// Setting up JSON headers
	w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode(animeTitles)
}
