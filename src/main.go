package main

import (
	"net/http"
	"os"

	"github.com/eoussama/anusic-api/src/shared/enums"
	hdlr "github.com/eoussama/anusic-api/src/shared/handlers"
	"github.com/eoussama/anusic-api/src/shared/middlewares"
	"github.com/eoussama/anusic-api/src/shared/scraper"
	"github.com/eoussama/anusic-api/src/shared/utils"
	v1 "github.com/eoussama/anusic-api/src/v1"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Loading environment variables
	utils.LoadEnvVars()

	// Loading cache data if available
	if !utils.LoadCache() {

		// Scraping if no data cached
		scraper.Scrap()

		// Saving the data
		utils.SaveCache(utils.Cache)
	}

	// Creating routers
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Use(middlewares.ContentType)

	apiRouter.HandleFunc("/", hdlr.IndexHandler).Methods("GET")
	apiRouter.HandleFunc("", hdlr.IndexHandler).Methods("GET")

	// Initializing the v1 API
	v1.Init(apiRouter)

	// CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})

	// Starting
	utils.Log("Starting...", enums.LogInfo)
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(corsObj)(router))
}
