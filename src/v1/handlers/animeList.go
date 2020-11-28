package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/eoussama/anusic-api/src/shared/utils"
// 	"github.com/eoussama/anusic-api/src/v1/models"
// 	"github.com/ulule/deepcopier"
// )

// // AnimeListHandler handles the anime list request (/api/v1/anime/)
// func AnimeListHandler(w http.ResponseWriter, r *http.Request) {

// 	// Scraping anime list
// 	animeTitlesEx := []models.AnimeEx{}
// 	animeTitles := utils.CachedAnimeList

// 	// If no cache available scrap data
// 	if len(animeTitles) == 0 {
// 		animeTitles = utils.ScrapAnimeList()
// 	}

// 	// Sanitizing the export struct
// 	for _, anm := range animeTitles {
// 		animeEx := models.AnimeEx{}

// 		deepcopier.Copy(&animeEx).From(anm)
// 		animeTitlesEx = append(animeTitlesEx, animeEx)
// 	}

// 	// Encoding the return value
// 	json.NewEncoder(w).Encode(animeTitlesEx)
// }
