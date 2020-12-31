package v1

import (
	hdlr "github.com/eoussama/anusic-api/src/v1/handlers"
	"github.com/eoussama/anusic-api/src/v1/handlers/anime"
	"github.com/eoussama/anusic-api/src/v1/handlers/log"
	"github.com/gorilla/mux"
)

// Init initializes the v1 of the API
func Init(r *mux.Router) {

	// Subrouting
	v1Route := r.PathPrefix("/v1").Subrouter()

	// Subroot routing
	v1Route.HandleFunc("/", hdlr.IndexHandler).Methods("GET")
	v1Route.HandleFunc("", hdlr.IndexHandler).Methods("GET")

	// Anime routing
	anime.Init(v1Route.PathPrefix("/anime").Subrouter())

	// Logs routing
	log.Init(v1Route.PathPrefix("/logs").Subrouter())
}
