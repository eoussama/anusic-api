package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

const base string = "https://www.reddit.com/r/AnimeThemes/wiki/"

func main() {
	r := mux.NewRouter()
	log.SetPrefix("[Anusic API] ")

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/anime", animeList).Methods("GET")

	log.Println("Starting...")
	http.ListenAndServe(":8000", r)
}

// Entry
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Anusic API")
}

// Get anime list
func animeList(w http.ResponseWriter, r *http.Request) {

	// Initializing the scraper
	collector := colly.NewCollector()

	// Initializing the anime list
	animeTitles := []Anime{}

	var currAnime string = ""

	// Scraping the catalog
	collector.OnHTML(".wiki", func(e *colly.HTMLElement) {

		// Anime list
		e.ForEach("#wiki_0-9 ~ p a", func(index int, element *colly.HTMLElement) {

			// Extracting anime title
			extract := element.Text
			idx := strings.LastIndex(extract, " (")
			year, _ := strconv.ParseInt(extract[idx+2:len(extract)-1], 10, 16)

			anime := Anime{
				Name:     extract[:idx],
				Year:     uint16(year),
				linkInfo: element.Attr("href"),
			}

			// Appending extracted anime title
			animeTitles = append(animeTitles, anime)

			currAnime = anime.linkInfo[strings.Index(anime.linkInfo, "#")+1:]
			collector.Visit("https://www.reddit.com" + anime.linkInfo)
		})

		// Anime info
		if len(currAnime) > 0 {

			// Getting all h3 elements
			e.ForEach("h3", func(_ int, element *colly.HTMLElement) {

				// Checking if the ID matches the current anime title
				if element.Attr("id") == currAnime {

					// Getting the target anime
					anime := &animeTitles[len(animeTitles)-1]

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
						anime.AltName = element.DOM.Next().Text()
					}

					fmt.Printf("%+v\n", animeTitles[len(animeTitles)-1])
				}
			})

			// Clearing the current anime
			currAnime = ""
		}
	})

	collector.OnRequest(func(req *colly.Request) {
		log.Println("Visiting " + req.URL.String())
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Fatal(err)
	})

	// Visiting the target page and invoking the scraper
	collector.Visit(base + "anime_index")

	// Setting up JSON headers
	w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode(animeTitles)
}
