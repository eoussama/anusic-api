package scraper

import (
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// Scrap scraps all data
func Scrap() {
	utils.Cache.Anime = AnimeList()
}
