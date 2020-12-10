package models

// Collection type, contains a slice of themes
type Collection struct {
	ID         string
	AnimeMALID uint16
	Name       string
}

// CollectionEx type
type CollectionEx struct {
	Name   string    `json:"name,omitempty"`
	Themes []ThemeEx `json:"themes"`
}

// FormatEx formats the struct
func (c Collection) FormatEx() CollectionEx {
	return CollectionEx{
		Name:   c.Name,
		Themes: []ThemeEx{},
	}
}
