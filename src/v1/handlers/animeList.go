package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eoussama/anusic-api/src/v1/utils"
)

// AnimeListHandler handles the anime list request (/api/v1/anime/)
func AnimeListHandler(w http.ResponseWriter, r *http.Request) {

	// Scraping anime list
	animeTitles := utils.CachedAnimeList

	// If no cache available scrap data
	if len(animeTitles) == 0 {
		animeTitles = utils.ScrapAnimeList()
	}

	// Encoding the return value
	json.NewEncoder(w).Encode(animeTitles)
}
