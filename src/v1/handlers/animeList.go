package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// AnimeListHandler handles the anime list request (/api/v1/anime/)
func AnimeListHandler(w http.ResponseWriter, r *http.Request) {

	// Scraping anime list
	animeTitles := []models.AnimeEx{}

	// Sanitizing the export struct
	for _, anime := range utils.Cache.Anime {
		animeTitles = append(animeTitles, anime.Format())
	}

	// Encoding the return value
	json.NewEncoder(w).Encode(animeTitles)
}
