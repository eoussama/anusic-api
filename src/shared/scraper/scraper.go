package scraper

import (
	"regexp"
	"strconv"
	"time"

	"github.com/eoussama/anusic-api/src/shared/enums"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// Scrap scraps all data
func Scrap() {
	AnimeList()
	AnimeInfo()
}

// genYears generates years from current to 60s
func genYears() []string {
	years := []string{}

	for year := time.Now().Year(); year >= 2000; year-- {
		years = append(years, strconv.Itoa(year))
	}

	return append(years, []string{"90s", "80s", "70s", "60s"}...)
}

// parseYear remove any characters from year and return an int
func parseYear(x string) uint16 {
	reg, err := regexp.Compile("[^0-9]*")

	if err != nil {
		utils.Log(err, enums.LogError)
	}

	ret, _ := strconv.ParseInt(reg.ReplaceAllString(x, ""), 10, 16)
	return uint16(ret)
}
