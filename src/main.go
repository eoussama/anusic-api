package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const base string = "https://www.reddit.com/r/AnimeThemes/wiki/"

func main() {
	r := mux.NewRouter()
	log.SetPrefix("[Anusic API] ")

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/anime", animeList).Methods("GET")

	log.Println("Starting...")
	// http.DefaultClient.Timeout = time.Minute * 10
	http.ListenAndServe(":8000", r)
}

// Entry
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Anusic API")
}

// Get anime list
func animeList(w http.ResponseWriter, r *http.Request) {
	// Scraping anime list
	animeTitles := scrapAnimeList()

	// Setting up JSON headers
	w.Header().Set("Content-Type", "application/json")

	// Encoding the return value
	json.NewEncoder(w).Encode(animeTitles)
}
