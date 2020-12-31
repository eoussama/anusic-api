package scraper

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/eoussama/anusic-api/src/shared/enums"
	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
	"github.com/gocolly/colly"
)

// AnimeList scraps the entire anime list
func AnimeList() {
	utils.Log("Scraping Anime list...", enums.LogInfo)
	start := time.Now()

	// Initializing the scraper
	collector := colly.NewCollector(colly.Async(true))

	// Initializing the anime list
	animeTitles := []models.Anime{}

	// Scraping the Anime list
	collector.OnHTML("#wiki_0-9 ~ p", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("a", func(_ int, element *colly.HTMLElement) bool {

			// Extracting anime title
			extract := element.Text
			idx := strings.LastIndex(extract, " (")
			year := extract[idx+2 : len(extract)-1]
			intYear := parseYear(year)

			anime := models.Anime{
				ID:   strings.ToLower(strings.Trim(strings.Replace(extract[:idx], " ", "", -1), " ")) + year,
				Name: strings.Trim(strings.Replace(extract[:idx], "\"", "", -1), " "),
				Year: intYear,
			}

			// Appending extracted anime title
			animeTitles = append(animeTitles, anime)
			return false
		})
	})

	// Visiting the target page and invoking the scraper
	collector.Visit(BASE_ANIME + "anime_index")
	collector.Wait()

	utils.Cache.Anime = animeTitles
	utils.Log(fmt.Sprintf("Fetched %d Anime titles in %v", len(animeTitles), time.Since(start)), enums.LogInfo)
}

// AnimeInfo scraps Anime info
func AnimeInfo() {
	utils.Log("Scraping Anime Info...", enums.LogInfo)

	start := time.Now()
	async := true
	count := 0
	years := genYears()

	// Initializing the scraper
	collector := colly.NewCollector(colly.Async(async))
	collector.SetRequestTimeout(60 * time.Second)

	if async {
		collector.Limit(&colly.LimitRule{
			DomainGlob:  "*",
			Parallelism: len(years),
		})
	}

	// Scraping the Anime info
	collector.OnHTML(".md.wiki > h3", func(e *colly.HTMLElement) {

		// Getting the Anime index
		year := path.Base(e.Request.URL.Path)
		targetID := strings.ToLower(strings.Trim(strings.Replace(e.Text, " ", "", -1), " ")) + year
		index, _ := utils.Cache.GetAnimeByID(targetID)

		if index > -1 {

			// Initializing the themes table selection
			var tableSelection *goquery.Selection

			// Getting the respective Anime
			anime := &utils.Cache.Anime[index]

			// Extracting the MAL ID
			mal := e.ChildAttr("a", "href")
			re := regexp.MustCompile("[0-9]+")
			res := re.FindAllString(mal, -1)

			if len(res) > 0 {
				id, _ := strconv.Atoi(res[0])
				anime.MALID = uint16(id)
			}

			// Extracting the alt name
			if e.DOM.Next().Is("p") && len(e.DOM.Next().Has("strong").Nodes) > 0 {
				altNamesStr := strings.Replace(e.DOM.Next().Text(), "\"", "", -1)
				altNamesFrg := strings.Split(altNamesStr, ",")
				anime.AltNames = []string{}

				for i := 0; i < len(altNamesFrg); i++ {
					anime.AltNames = append(anime.AltNames, altNamesFrg[i])
				}

				tableSelection = e.DOM.Next().Next()
			} else {
				tableSelection = e.DOM.Next()
			}

			// Scrapping themes
			Themes(anime.MALID, tableSelection)

			count++
		} else {
			utils.Log(fmt.Sprintf("Anime “%s” not found year = %s", targetID, year), enums.LogWarning)
		}
	})

	// Fallback on error
	collector.OnError(func(r *colly.Response, e error) {
		utils.Log(fmt.Sprintf("Re-requesting %s", r.Request.URL.String()), enums.LogInfo)
		collector.Visit(r.Request.URL.String())
	})

	for _, year := range years {

		// Constructing the year index page
		url := BASE_ANIME + year

		// Logging the request
		utils.Log(fmt.Sprintf("Requesting %s", url), enums.LogInfo)

		// Visiting the target page and invoking the scraper
		collector.Visit(url)

		// Sleeping for 1 second to avoid timing-out
		time.Sleep(1 * time.Second)
	}

	// Waiting for the scraping to resolve
	collector.Wait()

	// Logging the fetched status
	utils.Log(fmt.Sprintf("Fetched %d Anime info in %v", count, time.Since(start)), enums.LogInfo)

	// Raising a warning if the fetched info does not match the total Anime titles
	if count < len(utils.Cache.Anime) {
		utils.Log(fmt.Sprintf("Failed to fetch info of %d Anime title(s)", len(utils.Cache.Anime)-count), enums.LogWarning)
	}
}
