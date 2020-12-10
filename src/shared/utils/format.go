package utils

import (
	"github.com/eoussama/anusic-api/src/shared/models"
)

// FormatAnime returns a proper Anime object
func FormatAnime(anime models.Anime) models.AnimeEx {

	// Formating the anime
	formatedAnime := anime.FormatEx()

	collections := Cache.GetCollections(anime)
	themes := Cache.GetAnimeThemes(anime)

	// Getting the collections
	if len(collections) > 0 {
		for _, col := range collections {
			collection := col.FormatEx()
			formatedAnime.Collections = append(formatedAnime.Collections, collection)
		}
	} else {
		formatedAnime.Collections = []models.CollectionEx{{}}
	}

	// Getting the themes
	for index := range formatedAnime.Collections {
		for _, theme := range themes {
			if collections[index].ID == theme.CollectionID {
				formatedAnime.Collections[index].Themes = append(formatedAnime.Collections[index].Themes, theme.FormatEx())
			}
		}
	}

	return formatedAnime
}
