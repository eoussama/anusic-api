package models

import "fmt"

// Error type
type Error struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AnimeNotFound returns a proper error message for when a passed
// ID does not atch any scrapped Anime title
func (e Error) AnimeNotFound(id int) *Error {
	return &Error{
		Name:        "AnimeNotFound",
		Description: fmt.Sprintf("ID %d does not match any Anime title", id),
	}
}
