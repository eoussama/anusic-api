package scraper

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
	"github.com/gocolly/colly"
)

var wg sync.WaitGroup

// AnimeList scraps the entire anime list
func AnimeList() {
	log.Println("Scraping Anime list...")

	wg.Add(2)

	// Initializing the scraper
	collector := colly.NewCollector(
		colly.Async(true),
	)

	// Initializing the anime list
	animeTitles := []models.Anime{}

	// Scraping the catalog
	collector.OnHTML("#wiki_0-9 ~ p", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("a", func(index int, element *colly.HTMLElement) bool {

			// Extracting anime title
			extract := element.Text
			idx := strings.LastIndex(extract, " (")
			year, _ := strconv.ParseInt(extract[idx+2:len(extract)-1], 10, 16)
			href := element.Attr("href")
			targetID := href[strings.Index(href, "#")+1:]

			anime := models.Anime{
				ID:   targetID,
				Name: strings.Trim(strings.Replace(extract[:idx], "\"", "", -1), " "),
				Year: uint16(year),
			}

			// Element ID on the DOM
			if len(animeTitles) < 2 {
				go AnimeInfo(&anime)
			}

			// Appending extracted anime title
			animeTitles = append(animeTitles, anime)
			return true
		})
	})

	// Visiting the target page and invoking the scraper
	collector.Visit(os.Getenv("BASE") + "anime_index")

	collector.Wait()
	wg.Wait()

	utils.Cache.Anime = animeTitles
}

// AnimeInfo scraps Anime info
func AnimeInfo(anime *models.Anime) {

	collector := colly.NewCollector(
		colly.Async(true),
	)

	collector.OnHTML("h3", func(element *colly.HTMLElement) {
		if element.Attr("id") == anime.ID {

			// Extracting the ID
			mal := element.ChildAttr("a", "href")
			idx := strings.LastIndex(mal, "/anime/")
			extr := mal[idx+len("/anime/") : len(mal)-1]
			id, err := strconv.ParseInt(extr, 10, 32)

			if err == nil {
				anime.MALID = uint16(id)
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

			log.Printf("[%d] - %+v\n", 0, anime)
		}
	})

	collector.Visit(anime.GetLink())
	collector.Wait()

	defer wg.Done()
}
