package main

import (
	"log"
	"net/http"
	"os"

	hdlr "github.com/eoussama/anusic-api/src/shared/handlers"
	"github.com/eoussama/anusic-api/src/shared/middlewares"
	"github.com/eoussama/anusic-api/src/shared/scraper"
	"github.com/eoussama/anusic-api/src/shared/utils"
	v1 "github.com/eoussama/anusic-api/src/v1"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.SetPrefix("[Anusic API] ")

	// Loading environment variables
	utils.LoadEnvVars()

	// Loading cache data if available
	if !utils.LoadCache() {
		scraper.AnimeList()
	}

	// Creating routers
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	router.Use(middlewares.ContentType)
	router.HandleFunc("/", hdlr.IndexHandler).Methods("GET")

	// Initializing the v1 API
	v1.Init(apiRouter)

	// CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})

	// Starting
	log.Println("Starting...")
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(corsObj)(router))
}
