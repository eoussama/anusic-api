package main

import (
	"encoding/json"
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
	animeCol := []Anime{}

	// Scraping the catalog
	collector.OnHTML(".wiki", func(e *colly.HTMLElement) {
		e.ForEach("#wiki_0-9 ~ p a", func(_ int, element *colly.HTMLElement) {
			extract := element.Text
			idx := strings.LastIndex(extract, " (")
			year, _ := strconv.ParseInt(extract[idx+2:len(extract)-1], 10, 16)

			anime := Anime{
				Name: extract[:idx],
				Year: uint16(year),
			}

			animeCol = append(animeCol, anime)
		})
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
	json.NewEncoder(w).Encode(animeCol)
}
