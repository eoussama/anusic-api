package scraper

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
	"github.com/gocolly/colly"
)

// AnimeList scraps the entire anime list
func AnimeList() {
	log.Println("Scraping Anime list...")

	// Initializing the scraper
	collector := colly.NewCollector(colly.Async(true))

	// Initializing the anime list
	animeTitles := []models.Anime{}

	// Scraping the catalog
	collector.OnHTML("#wiki_0-9 ~ p", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("a", func(_ int, element *colly.HTMLElement) bool {

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

			// Appending extracted anime title
			animeTitles = append(animeTitles, anime)
			return false
		})
	})

	// Visiting the target page and invoking the scraper
	collector.Visit(os.Getenv("BASE") + "anime_index")
	collector.Wait()

	utils.Cache.Anime = animeTitles
}

// AnimeInfo scraps Anime info
func AnimeInfo() {
	log.Println("Scraping Anime Info...")
	start := time.Now()

	collector := colly.NewCollector(colly.Async(true))

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})

	collector.OnHTML(".md.wiki", func(e *colly.HTMLElement) {

		// Getting the current Anime
		url := e.Request.URL.String()
		index, _ := strconv.Atoi(url[strings.LastIndex(url, "?index=")+7:])
		anime := &utils.Cache.Anime[index]

		e.ForEachWithBreak("h3", func(_ int, element *colly.HTMLElement) bool {
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

				log.Printf("[%d] - %+v\n", index, anime)
				return false
			}

			return true
		})
	})

	for index, anime := range utils.Cache.Anime[:50] {
		if index < 100 {
			collector.Visit(anime.GetLink() + "?index=" + strconv.Itoa(index))
		}
	}

	// collector.Wait()
	fmt.Println("Elapsed time: ", time.Since(start))
}
