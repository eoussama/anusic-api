package utils

import (
	"github.com/eoussama/anusic-api/src/shared/models"
)

// FormatAnime returns a proper Anime object
func FormatAnime(anime models.Anime) models.AnimeEx {

	// Formating the anime
	formatedAnime := anime.Format()

	for _, theme := range Cache.GetAnimeThemes(anime) {
		formatedAnime.Themes = append(formatedAnime.Themes, theme)
	}
	return formatedAnime
}
