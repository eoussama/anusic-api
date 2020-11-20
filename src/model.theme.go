package main

// Theme song type
type Theme struct {
	Name         string   `json:"name"`
	Link         string   `json:"link"`
	Episodes     []string `json:"episodes"`
	ThemeType    int8     `json:"type"`
	Format       int8     `json:"format"`
	Resolution   string   `json:"resolution"`
	HasSpoilers  bool     `json:"hasSpoilers"`
	IsNSFW       bool     `json:"isNSFW"`
	IsCreditless bool     `json:"isCreditless"`
	HasLyrics    bool     `json:"hasLyrics"`
	IsTransition bool     `json:"isTransition"`
	IsOver       bool     `json:"isOver"`
}
