package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.SetPrefix("[Anusic API] ")

	// Loading environment variables
	loadEnvVars()

	// Loading cache data if available
	loadCache()

	// Creating router
	router := mux.NewRouter()

	// Routing
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/anime", AnimeListHandler).Methods("GET")
	router.HandleFunc("/anime/{id:[0-9]+}", AnimeHandler).Methods("GET")

	// CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})

	// Starting
	log.Println("Starting...")
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(corsObj)(router))
}
