package models

// Source type
type Source struct {
	Link         string
	Format       string
	Resolution   string
	HasLyrics    bool
	IsTransition bool
	IsOver       bool
	Tags         []string
}

// SourceEx export type
type SourceEx struct {
	Link         string   `json:"link"`
	Format       string   `json:"format"`
	Resolution   string   `json:"resolution"`
	HasLyrics    bool     `json:"hasLyrics"`
	IsTransition bool     `json:"isTransition"`
	IsOver       bool     `json:"isOver"`
	Tags         []string `json:"tags,omitempty"`
}

// FormatEx formats the struct
func (s Source) FormatEx() SourceEx {
	return SourceEx{
		Link:         s.Link,
		Format:       s.Format,
		Resolution:   s.Resolution,
		HasLyrics:    s.HasLyrics,
		IsTransition: s.IsTransition,
		IsOver:       s.IsOver,
		Tags:         s.Tags,
	}
}
