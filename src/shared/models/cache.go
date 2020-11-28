package models

// Cache type used to store cached anime titles
type Cache struct {
	Anime  []Anime
	Themes []Theme
}

// GetAnime gets an Anime by ID
func (c Cache) GetAnime(id int) *Anime {
	if len(c.Anime) > 0 {
		for _, anime := range c.Anime {
			if anime.MALID == uint16(id) {
				return &anime
			}
		}
	}

	return nil
}
