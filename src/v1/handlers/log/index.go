package log

import (
	"github.com/gorilla/mux"
)

// Init initializes the logs subroute
func Init(r *mux.Router) {

	// Logs
	r.HandleFunc("", Logs).Methods("GET")
	r.HandleFunc("/", Logs).Methods("GET")

	// Log by ID
	r.HandleFunc("/{id}", Log).Methods("GET")
	r.HandleFunc("/{id}/", Log).Methods("GET")
}
