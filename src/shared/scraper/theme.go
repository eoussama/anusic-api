package scraper

import (
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

			// Theme title
			if i == 0 {
				var reg *regexp.Regexp
				dump := s.Text()
				fragments := getTitleFragments(dump)

				// Extracting the title
				reg = regexp.MustCompile(`".*"`)
				themeTitle := reg.FindStringSubmatch(dump)

				if len(themeTitle) > 0 {
					theme.Name = themeTitle[0][1 : len(themeTitle[0])-1]
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
				reg = regexp.MustCompile(`[0-9]*`)
				themeOrder := reg.FindStringSubmatch(fragments[0])

				if len(themeOrder) > 0 {
					order, _ := strconv.Atoi(themeOrder[0])
					theme.Order = uint8(order)
				} else {
					theme.Order = 1
				}
			}
		})

		// fmt.Printf("%+v\n", theme)
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
