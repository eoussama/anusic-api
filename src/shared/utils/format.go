package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eoussama/anusic-api/src/shared/enums"
	"github.com/eoussama/anusic-api/src/shared/models"
)

// FormatAnime returns a proper Anime object
func FormatAnime(anime models.Anime) models.AnimeEx {

	// Formating the anime
	formatedAnime := anime.FormatEx()

	associatedCols := Cache.GetCollections(anime)
	formatedCols := []models.CollectionEx{}
	themes := Cache.GetAnimeThemes(anime)

	// Getting the collections
	if len(associatedCols) > 0 {
		for _, col := range associatedCols {
			collection := col.FormatEx()
			formatedCols = append(formatedCols, collection)
		}
	}

	// Getting the themes
	for index := range formatedCols {
		for _, theme := range themes {
			if associatedCols[index].ID == theme.CollectionID {
				formatedCols[index].Themes = append(formatedCols[index].Themes, theme.FormatEx())
			}
		}

		if len(formatedCols[index].Themes) > 0 {
			formatedAnime.Collections = append(formatedAnime.Collections, formatedCols[index])
		}
	}

	return formatedAnime
}

// ReturnResponse returns the response object
func ReturnResponse(w http.ResponseWriter, r *http.Request, data interface{}, err *models.Error, status int) {

	// Logging the request
	Log(fmt.Sprintf("Path(%s), Queries(%s)", r.URL.Path, r.URL.RawQuery), enums.LogRequest)

	// Header
	w.WriteHeader(status)

	// Returning the marshalled data
	json.NewEncoder(w).Encode(models.Response{
		Status:   status,
		HasError: err != nil,
		Error:    err,
		Data:     data,
	})
}
