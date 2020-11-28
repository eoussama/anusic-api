package models

import (
	"os"
	"strconv"
)

// Anime type
type Anime struct {
	ID       string
	MALID    uint16
	Name     string
	AltNames []string
	Year     uint16
	Themes   []Theme
}

// AnimeEx export type
type AnimeEx struct {
	ID       uint16   `json:"id"`
	Name     string   `json:"name"`
	AltNames []string `json:"altNames"`
	Year     uint16   `json:"year"`
	Themes   []Theme  `json:"themes"`
}

// GetLink constructs the Anime info link
func (a Anime) GetLink() string {
	return os.Getenv("BASE") + strconv.Itoa(int(a.Year)) + "#" + a.ID
}

// JSON formats the struct
func (a Anime) JSON() AnimeEx {
	return AnimeEx{
		ID:       a.MALID,
		Name:     a.Name,
		AltNames: a.AltNames,
		Year:     a.Year,
		Themes:   a.Themes,
	}
}
