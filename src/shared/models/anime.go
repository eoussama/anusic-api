package models

import (
	"regexp"
	"strconv"
	"strings"
)

// Anime type
type Anime struct {
	ID       string
	MALID    uint16
	Name     string
	AltNames []string
	Year     uint16
}

// AnimeEx export type
type AnimeEx struct {
	ID          uint16         `json:"id"`
	Name        string         `json:"name"`
	AltNames    []string       `json:"altNames,omitempty"`
	Year        uint16         `json:"year"`
	Collections []CollectionEx `json:"collections,omitempty"`
}

// MatchName Checks if string matches any of the ANime title's name
func (a Anime) MatchName(name string) bool {

	// Trimming the name
	mName := strings.Trim(name, " ")

	// Returning true if no name was passed
	if len(name) == 0 {
		return true
	}

	// Building the regular expression
	reg := "(?i).*"
	for _, frag := range strings.Split(mName, " ") {
		reg += frag + ".*"
	}
	exp := regexp.MustCompile(reg)

	// Iterating over the name and alt names
	for _, animeName := range append([]string{a.Name}, a.AltNames...) {
		if exp.MatchString(animeName) {
			return true
		}
	}

	// Returning false if no matches were found
	return false
}

// MatchYear Checks if the anime belongs to any of the input years' releases
func (a Anime) MatchYear(year string) bool {

	for _, y := range strings.Split(year, ",") {
		sanYear, _ := strconv.Atoi(y)

		if a.Year == uint16(sanYear) || len(year) == 0 {
			return true
		}
	}

	return false
}

// FormatEx formats the struct
func (a Anime) FormatEx() AnimeEx {
	return AnimeEx{
		ID:          a.MALID,
		Name:        a.Name,
		AltNames:    a.AltNames,
		Year:        a.Year,
		Collections: []CollectionEx{},
	}
}
