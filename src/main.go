package main

import (
	"log"
	"net/http"
	"os"

	hdlr "github.com/eoussama/anusic-api/src/handlers"
	"github.com/eoussama/anusic-api/src/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.SetPrefix("[Anusic API] ")

	// Loading environment variables
	utils.LoadEnvVars()

	// Loading cache data if available
	utils.LoadCache()

	// Creating routers
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	// Routing
	router.HandleFunc("/", hdlr.IndexHandler).Methods("GET")

	apiRouter.HandleFunc("/anime", hdlr.AnimeListHandler).Methods("GET")
	apiRouter.HandleFunc("/anime/", hdlr.AnimeListHandler).Methods("GET")
	apiRouter.HandleFunc("/anime/{id:[0-9]+}", hdlr.AnimeHandler).Methods("GET")

	// CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})

	// Starting
	log.Println("Starting...")
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(corsObj)(router))
}
