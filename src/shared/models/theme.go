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
	// Link         string
	// Format       int8
	// Resolution   string
	// IsCreditless bool
	// HasLyrics    bool
	// IsTransition bool
	// IsOver       bool
}
