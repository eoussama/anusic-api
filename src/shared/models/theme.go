package models

// Theme song type
type Theme struct {
	Name         string
	Link         string
	Episodes     []string
	ThemeType    int8
	Format       int8
	Resolution   string
	HasSpoilers  bool
	IsNSFW       bool
	IsCreditless bool
	HasLyrics    bool
	IsTransition bool
	IsOver       bool
}
