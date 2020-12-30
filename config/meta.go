package config

// Meta represents the structure of the inf object of the API
type Meta struct {
	Name       string
	Author     string
	Repository string
}

// Info information about the API
var Info = Meta{
	"Anusic API",
	"EOussama",
	"https://github.com/EOussama/anusic-api",
}
