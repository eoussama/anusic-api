package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/anime", animeList).Methods("GET")

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
	c := colly.NewCollector()

	// Catalog list (used to group target entries)
	catalog := []string{}

	// Initializing the anime list
	// animeCol := []string{}

	// Scraping the catalog
	c.OnHTML(".toc", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, element *colly.HTMLElement) {
			catalog = append(catalog, element.Text)
		})
	})

	fmt.Println(catalog)

	// Visiting the target page and invoking the scraper
	c.Visit("https://www.reddit.com/r/AnimeThemes/wiki/anime_index")

	// Setting up JSON headers
	w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode("Anime list")
}

// https://www.reddit.com/r/AnimeThemes/wiki
// "github.com/gocolly/colly"
