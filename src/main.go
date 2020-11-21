package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const base string = "https://www.reddit.com/r/AnimeThemes/wiki/"

func main() {

	// Routing
	router := mux.NewRouter()
	log.SetPrefix("[Anusic API] ")

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/anime", animeList).Methods("GET")

	// Loading cache data if available
	loadCache()

	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Println("Starting...")
	http.ListenAndServe(":8000", handlers.CORS(corsObj)(router))
}

// Entry
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Anusic API")
}

// Get anime list
func animeList(w http.ResponseWriter, r *http.Request) {

	// Scraping anime list
	animeTitles := cachedAnimeList

	// If no cache available scrap data
	if len(animeTitles) == 0 {
		animeTitles = scrapAnimeList()
	}

	// Setting up JSON headers
	w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode(animeTitles)
}
