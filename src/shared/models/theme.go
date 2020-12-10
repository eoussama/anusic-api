package models

// Theme song type
type Theme struct {
	AnimeMALID   uint16
	CollectionID string
	Name         string
	ThemeType    uint8
	Order        uint8
	Episodes     []string
	IsNSFW       bool
	HasSpoilers  bool
	Sources      []Source
}

// ThemeEx export type
type ThemeEx struct {
	Name        string   `json:"name"`
	ThemeType   uint8    `json:"type"`
	Order       uint8    `json:"order"`
	Episodes    []string `json:"episodes"`
	IsNSFW      bool     `json:"isNSFW"`
	HasSpoilers bool     `json:"hasSpoilers"`
	Sources     []Source `json:"sources,omitempty"`
}

// Format formats the struct
func (a Theme) Format() ThemeEx {
	return ThemeEx{
		Name:        a.Name,
		ThemeType:   a.ThemeType,
		Order:       a.Order,
		Episodes:    a.Episodes,
		IsNSFW:      a.IsNSFW,
		HasSpoilers: a.HasSpoilers,
		Sources:     a.Sources,
	}
}
