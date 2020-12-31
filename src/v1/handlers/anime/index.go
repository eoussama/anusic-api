package anime

import (
	"github.com/gorilla/mux"
)

// Init initializes the anime subroute
func Init(r *mux.Router) {

	// Anime list
	r.HandleFunc("", AnimeListHandler).Methods("GET")
	r.HandleFunc("/", AnimeListHandler).Methods("GET")

	// Anime by ID
	r.HandleFunc("/{id:[0-9]+}", AnimeHandler).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}/", AnimeHandler).Methods("GET")
}
