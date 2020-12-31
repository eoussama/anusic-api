package anime

import (
	"net/http"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// AnimeListHandler handles the anime list request (/api/v1/anime/)
func AnimeListHandler(w http.ResponseWriter, r *http.Request) {

	// Getting the name query
	qName := r.URL.Query().Get("name")
	qYear := r.URL.Query().Get("year")

	// Export Anime list
	animeTitles := []models.AnimeEx{}

	// Sanitizing the export struct
	for _, anime := range utils.Cache.FilterAnime(qName, qYear) {
		animeTitles = append(animeTitles, anime.FormatEx())
	}

	// Returning the response
	utils.ReturnResponse(w, r, animeTitles, nil, http.StatusOK)
}
