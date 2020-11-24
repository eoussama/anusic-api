package v1

import (
	hdlr "github.com/eoussama/anusic-api/src/v1/handlers"
	"github.com/gorilla/mux"
)

// Init initializes the v1 of the API
func Init(r *mux.Router) {

	// Routing
	v1Route := r.PathPrefix("/v1/").Subrouter()

	// Root
	v1Route.HandleFunc("/", hdlr.IndexHandler).Methods("GET")

	// Anime list
	v1Route.HandleFunc("/anime", hdlr.AnimeListHandler).Methods("GET")
	v1Route.HandleFunc("/anime/", hdlr.AnimeListHandler).Methods("GET")

	// Anime by ID
	v1Route.HandleFunc("/anime/{id:[0-9]+}", hdlr.AnimeHandler).Methods("GET")
}
