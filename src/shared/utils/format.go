package utils

import (
	"encoding/json"
	"net/http"

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

// FormatResponse formats the output response
func FormatResponse(data interface{}, err *models.Response) map[string]interface{} {
	output := map[string]interface{}{}
	response := err

	if err == nil {
		response = &models.Response{}
	}

	output["hasError"] = response.HasError
	output["error"] = response.Error
	output["data"] = data

	return output
}

// ReturnResponse returns the response object
func ReturnResponse(w http.ResponseWriter, data interface{}) {
	response := FormatResponse(data, nil)
	json.NewEncoder(w).Encode(response)
}
