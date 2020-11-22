package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const base string = "https://www.reddit.com/r/AnimeThemes/wiki/"

func main() {
	log.SetPrefix("[Anusic API] ")

	// Loading environment variables
	loadEnvVars()

	// Loading cache data if available
	loadCache()

	// Routing
	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/anime", AnimeListHandler).Methods("GET")

	log.Println("Starting...")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(corsObj)(router))
}
