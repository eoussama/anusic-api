package models

// AnimeEx export type
type AnimeEx struct {
	ID       uint16   `json:"id"`
	Name     string   `json:"name"`
	AltNames []string `json:"altNames"`
	Year     uint16   `json:"year"`
	Themes   []Theme  `json:"themes"`
}
