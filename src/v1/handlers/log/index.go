package log

import (
	"github.com/eoussama/anusic-api/src/shared/middlewares"
	"github.com/gorilla/mux"
)

// Init initializes the logs subroute
func Init(r *mux.Router) {

	// Invoking the auth middleware
	r.Use(middlewares.Auth)

	// Logs
	r.HandleFunc("", Logs).Methods("GET")
	r.HandleFunc("/", Logs).Methods("GET")

	// Log by ID
	r.HandleFunc("/{id}", Log).Methods("GET")
	r.HandleFunc("/{id}/", Log).Methods("GET")
}
