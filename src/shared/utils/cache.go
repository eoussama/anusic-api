package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/eoussama/anusic-api/src/shared/models"
)

// Cache caches all scraped data
var Cache models.Cache = models.Cache{}

// LoadCache loads the cache files
func LoadCache() bool {
	log.Println("Loading cache...")

	// Constructing the cache file
	absPath, _ := filepath.Abs(".")
	path := filepath.Join(absPath, "data", "cachee.json")
	jsonFile, err := os.Open(path)

	// No cache found
	if err != nil {
		return false
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Cache)

	// Clearing resources
	defer jsonFile.Close()
	return true
}
