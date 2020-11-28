package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/eoussama/anusic-api/src/shared/utils"
// 	"github.com/eoussama/anusic-api/src/v1/models"
// 	"github.com/ulule/deepcopier"

// 	"github.com/gorilla/mux"
// )

// // AnimeHandler handles the anime request (/api/v1/anime/{id:[0-9]+})
// func AnimeHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	// Scraping anime list
// 	anime := models.Anime{}
// 	animeEx := models.AnimeEx{}

// 	// If no cache available scrap data
// 	if len(utils.CachedAnimeList) > 0 {
// 		for _, anm := range utils.CachedAnimeList {
// 			if anm.ID == uint16(id) {
// 				anime = anm
// 			}
// 		}
// 	}

// 	// Sanitizing the export struct
// 	deepcopier.Copy(&animeEx).From(anime)

// 	// Encoding the return value
// 	json.NewEncoder(w).Encode(animeEx)
// }
