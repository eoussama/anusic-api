package models

// Cache type used to store cached anime titles
type Cache struct {
	Anime       []Anime
	Themes      []Theme
	Collections []Collection
}

// GetAnimeByMALID gets an Anime by MAL ID
func (c Cache) GetAnimeByMALID(id int) *Anime {
	if len(c.Anime) > 0 {
		for _, anime := range c.Anime {
			if anime.MALID == uint16(id) {
				return &anime
			}
		}
	}

	return nil
}

// GetAnimeByID gets an Anime by ID
func (c Cache) GetAnimeByID(id string) (int, *Anime) {
	if len(c.Anime) > 0 {
		for index, anime := range c.Anime {
			if anime.ID == id {
				return index, &anime
			}
		}
	}

	return -1, nil
}

// GetAnimeThemes gets all themes associated with an Anime
func (c Cache) GetAnimeThemes(anime Anime) []Theme {
	var retrievedThemes []Theme = []Theme{}

	if len(c.Themes) > 0 {
		for _, theme := range c.Themes {
			if theme.AnimeMALID == anime.MALID {
				retrievedThemes = append(retrievedThemes, theme)
			}
		}
	}

	return retrievedThemes
}

// GetCollections gets associated collections
func (c Cache) GetCollections(anime Anime) []Collection {
	var retrievedCollections []Collection = []Collection{}

	if len(c.Collections) > 0 {
		for _, collection := range c.Collections {
			if collection.AnimeMALID == anime.MALID {
				retrievedCollections = append(retrievedCollections, collection)
			}
		}
	}

	return retrievedCollections
}
