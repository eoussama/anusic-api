package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Media format
const (
	MediaFormatWebm = iota
)

// Theme song types
const (
	ThemeTypeOP = iota
	ThemeTypeED = iota
)

// Theme song type
type Theme struct {
	name         string
	link         string
	episodes     []string
	themeType    int8
	format       int8
	resolution   string
	hasSpoilers  bool
	isNSFW       bool
	isCreditless bool
	hasLyrics    bool
	isTransition bool
	isOver       bool
}

// Anime type
type Anime struct {
	id      string
	name    string
	altName string
	themes  []Theme
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)

	http.ListenAndServe(":8000", r)
}

func index(w http.ResponseWriter, r *http.Request) {

}
