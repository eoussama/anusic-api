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
	path := filepath.Join(absPath, "data", "cache.json")
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

// SaveCache saves the data
func SaveCache(cache models.Cache) {
	log.Println("Saving cache...")

	// Constructing the cache file
	absPath, _ := filepath.Abs(".")
	path := filepath.Join(absPath, "data", "cache.json")

	// Marshalling the data
	file, _ := json.MarshalIndent(cache, "", " ")

	// Writing to the export file
	_ = os.Mkdir("data", 0755)
	_ = ioutil.WriteFile(path, file, 0644)

	log.Println("Cache saved in " + path)
}
