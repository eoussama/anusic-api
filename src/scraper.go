package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func scrapAnimeList() []Anime {

	// Initializing the scraper
	collector := colly.NewCollector(
		colly.Async(true),
	)

	// Initializing the anime list
	animeTitles := []Anime{}

	// Scraping the catalog
	collector.OnHTML("#wiki_0-9 ~ p", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("a", func(_ int, element *colly.HTMLElement) bool {

			// Extracting anime title
			extract := element.Text
			idx := strings.LastIndex(extract, " (")
			year, _ := strconv.ParseInt(extract[idx+2:len(extract)-1], 10, 16)

			anime := Anime{
				Name:     strings.Trim(strings.Replace(extract[:idx], "\"", "", -1), " "),
				Year:     uint16(year),
				linkInfo: element.Attr("href"),
			}

			// Element ID on the DOM
			targetID := anime.linkInfo[strings.Index(anime.linkInfo, "#")+1:]
			scrapAnimeInfo(targetID, &anime)
			log.Printf("[%d] - %+v\n", len(animeTitles), anime)

			// Appending extracted anime title
			animeTitles = append(animeTitles, anime)
			return true
		})
	})

	// Visiting the target page and invoking the scraper
	collector.Visit(os.Getenv("BASE") + "anime_index")
	collector.Wait()

	return animeTitles
}

func scrapAnimeInfo(targetID string, anime *Anime) {

	collector := colly.NewCollector(
		colly.Async(true),
	)

	collector.OnHTML("h3", func(element *colly.HTMLElement) {
		if element.Attr("id") == targetID {

			// Extracting the ID
			mal := element.ChildAttr("a", "href")
			idx := strings.LastIndex(mal, "/anime/")
			extr := mal[idx+len("/anime/") : len(mal)-1]
			id, err := strconv.ParseInt(extr, 10, 32)

			if err == nil {
				anime.ID = uint16(id)
			}

			// Extracting the alt name
			if element.DOM.Next().Is("p") {
				altNamesStr := strings.Replace(element.DOM.Next().Text(), "\"", "", -1)
				altNamesFrg := strings.Split(altNamesStr, ",")
				anime.AltNames = []string{}

				for i := 0; i < len(altNamesFrg); i++ {
					anime.AltNames = append(anime.AltNames, altNamesFrg[i])
				}
			}
		}
	})

	collector.Visit("https://www.reddit.com" + anime.linkInfo)
	collector.Wait()
}
