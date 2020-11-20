package main

// Anime type
type Anime struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	AltName string  `json:"altName"`
	Year    uint16  `json:"year"`
	Themes  []Theme `json:"themes"`
}
