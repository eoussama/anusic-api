package config

// Meta represents the structure of the inf object of the API
type Meta struct {
	Name       string `json:"name"`
	Author     string `json:"author"`
	Repository string `json:"repository"`
}

// Info information about the API
var Info = Meta{
	"Anusic API",
	"EOussama",
	"https://github.com/EOussama/anusic-api",
}
