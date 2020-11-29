package scraper

import (
	"strconv"
	"time"
)

// Scrap scraps all data
func Scrap() {
	AnimeList()
	AnimeInfo()
}

// Generate years
func genYears() []string {
	years := []string{}

	for year := time.Now().Year(); year >= 2000; year-- {
		years = append(years, strconv.Itoa(year))
	}

	return append(years, []string{"90s", "80s", "70s", "60s"}...)
}
