package models

// Theme song type
type Theme struct {
	AnimeMALID  uint16
	Name        string
	ThemeType   uint8
	Order       uint8
	Episodes    []string
	IsNSFW      bool
	HasSpoilers bool
	Sources     []Source
}
