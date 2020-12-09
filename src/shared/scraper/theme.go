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
	var lastThemeIndex int
	var selection *goquery.Selection = e

	for selection.Is("table") || selection.Is("p") {
		var collection models.Collection

		// Checking for collection availability
		if !selection.Is("table") {

			// Constructing the collection's ID
			collection.ID = fmt.Sprintf("%s-%s", strconv.Itoa(int(malID)), strings.ToLower(strings.Trim(strings.Replace(selection.Text(), " ", "", -1), " ")))

			// Extracting the collection's name
			collection.Name = selection.Text()

			// Assign the sibling table
			selection = selection.Next()
		}

		// Extracting theme info
		selection.ChildrenFiltered("tbody").Children().Each(func(_ int, s *goquery.Selection) {
			sourceRow := false
			theme := models.Theme{}

			theme.AnimeMALID = malID
			theme.CollectionID = collection.ID

			s.Children().Each(func(i int, s *goquery.Selection) {
				var reg *regexp.Regexp
				var dump string = s.Text()

				// Theme title
				if i == 0 {

					// Checking if the row is a follow up link
					if len(dump) == 0 {
						sourceRow = true
					}

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

				// Links
				if i == 1 {
					source := models.Source{}
					reg = regexp.MustCompile(`\((.*)\)`)
					fragments := reg.FindStringSubmatch(dump)

					// Extracting the link
					source.Link, _ = s.Children().Attr("href")

					// Extracting the format
					reg = regexp.MustCompile(`([\w]*)`)
					format := reg.FindStringSubmatch(dump)

					if len(format) > 1 {
						source.Format = strings.ToLower(format[1])
					}

					// Extracting the resolution
					if len(fragments) > 1 {
						reg = regexp.MustCompile(`\d+`)
						res := reg.FindStringSubmatch(fragments[0])

						if len(res) > 0 {
							source.Resolution = res[0]
						}
					}

					// Extracting the tags
					if len(fragments) > 1 {
						source.Tags = getSourceTags(fragments[1], source)
					}

					if len(fragments) > 1 {

						// Extracting the lyrics status
						source.HasLyrics = strings.Contains(strings.ToLower(fragments[1]), strings.ToLower("lyrics"))

						// Extracting the over status
						source.IsOver = strings.Contains(strings.ToLower(fragments[1]), strings.ToLower("over"))

						// Extracting the trans status
						source.IsTransition = strings.Contains(strings.ToLower(fragments[1]), strings.ToLower("trans"))

					}

					// Adding the theme to the Anime title
					if !sourceRow {
						theme.Sources = append(theme.Sources, source)
					} else {
						utils.Cache.Themes[lastThemeIndex].Sources = append(utils.Cache.Themes[lastThemeIndex].Sources, source)
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

			if !sourceRow {
				utils.Cache.Themes = append(utils.Cache.Themes, theme)
				lastThemeIndex = len(utils.Cache.Themes) - 1
			}
		})

		// Saving the collection if it's valid
		if len(collection.ID) > 0 {
			utils.Cache.Collections = append(utils.Cache.Collections, collection)
		}

		// Getting the next sibling element
		selection = selection.Next()
	}
}

// Gets sanitized fragments from the title dump
func getTitleFragments(input string) []string {
	var ret []string

	for _, frag := range strings.Split(input, " ") {
		if !strings.Contains(frag, "\"") {
			ret = append(ret, frag)
		}
	}

	return ret
}

// Splits and trims a string
func sanitizedSplit(input string) []string {
	var ret []string

	for _, frag := range strings.Split(input, ",") {
		ret = append(ret, strings.TrimSpace(frag))
	}

	return ret
}

// Extracts the source tags
func getSourceTags(input string, source models.Source) []string {
	var ret []string = []string{}

	for _, frag := range strings.Split(input, ",") {
		sanFrag := strings.TrimSpace(strings.ToLower(frag))

		if sanFrag != "lyrics" && sanFrag != "trans" && sanFrag != "over" && sanFrag != source.Resolution {
			ret = append(ret, strings.TrimSpace(frag))
		}
	}

	return ret
}
