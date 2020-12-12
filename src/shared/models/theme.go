package models

// Theme song type
type Theme struct {
	AnimeMALID   uint16
	CollectionID string
	Name         string
	ThemeType    uint8
	Order        uint8
	Version      uint8
	Episodes     []string
	IsNSFW       bool
	HasSpoilers  bool
	Sources      []Source
}

// ThemeEx export type
type ThemeEx struct {
	Name        string     `json:"name"`
	ThemeType   uint8      `json:"type"`
	Order       uint8      `json:"order"`
	Variant     uint8      `json:"version"`
	Episodes    []string   `json:"episodes"`
	IsNSFW      bool       `json:"isNSFW"`
	HasSpoilers bool       `json:"hasSpoilers"`
	Sources     []SourceEx `json:"sources,omitempty"`
}

// FormatEx formats the struct
func (t Theme) FormatEx() ThemeEx {
	sources := []SourceEx{}

	// Formating the sources
	for _, source := range t.Sources {
		sources = append(sources, source.FormatEx())
	}

	return ThemeEx{
		Name:        t.Name,
		ThemeType:   t.ThemeType,
		Order:       t.Order,
		Variant:     t.Version,
		Episodes:    t.Episodes,
		IsNSFW:      t.IsNSFW,
		HasSpoilers: t.HasSpoilers,
		Sources:     sources,
	}
}
