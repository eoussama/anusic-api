package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")

	http.ListenAndServe(":8000", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	themes := []Theme{}
	themes = append(themes, Theme{
		Name: "dddd",
	})

	anime := Anime{
		Name:    ".hack//Liminality",
		AltName: "",
		ID:      "",
		Year:    2002,
		Themes:  themes,
	}

	fmt.Println(anime)
	json.NewEncoder(w).Encode(anime)
}
