package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/eoussama/anusic-api/src/shared/enums"
	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// Themes scraps related theme songs
func Themes(malID uint16, e *goquery.Selection) {
	e.ChildrenFiltered("tbody").Children().Each(func(_ int, s *goquery.Selection) {
		theme := models.Theme{}
		theme.AnimeMALID = malID

		s.Children().Each(func(i int, s *goquery.Selection) {
			var reg *regexp.Regexp
			var dump string = s.Text()

			// Theme title
			if i == 0 {
				fragments := getTitleFragments(dump)

				// Extracting the title
				reg = regexp.MustCompile(`"(.*)"`)
				themeTitle := reg.FindStringSubmatch(dump)

				if len(themeTitle) > 1 {
					theme.Name = themeTitle[1]
				}

				// Extracting the type
				reg = regexp.MustCompile(`[A-Z]*`)
				themeType := reg.FindStringSubmatch(fragments[0])

				if len(themeType) > 0 {
					if themeType[0] == "OP" {
						theme.ThemeType = enums.ThemeTypeOP
					}

					if themeType[0] == "ED" {
						theme.ThemeType = enums.ThemeTypeED
					}
				}

				// Extracting the order
				reg = regexp.MustCompile(`[0-9]+`)
				themeOrder := reg.FindStringSubmatch(fragments[0])

				if len(themeOrder) > 0 {
					order, _ := strconv.Atoi(themeOrder[0])
					theme.Order = uint8(order)
				} else {
					theme.Order = 1
				}
			}

			// Episodes
			if i == 2 {

				// Extracting the episodes
				theme.Episodes = sanitizedSplit(dump)
			}

			// Notes
			if i == 3 {

				// Extracting NSFW status
				theme.IsNSFW = strings.Contains(strings.ToLower(dump), strings.ToLower("nsfw"))

				// Extracting spoilers status
				theme.HasSpoilers = strings.Contains(strings.ToLower(dump), strings.ToLower("spoiler"))
			}
		})

		fmt.Printf("%+v\n", theme)
		utils.Cache.Themes = append(utils.Cache.Themes, theme)
	})
}

// Gets sanitized fragments from the title dump
func getTitleFragments(dump string) []string {
	var ret []string

	for _, frag := range strings.Split(dump, " ") {
		if !strings.Contains(frag, "\"") {
			ret = append(ret, frag)
		}
	}

	return ret
}

// Splits and trims a string
func sanitizedSplit(dump string) []string {
	var ret []string

	for _, frag := range strings.Split(dump, ",") {
		ret = append(ret, strings.TrimSpace(frag))
	}

	return ret
}
