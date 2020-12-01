package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/eoussama/anusic-api/src/shared/enums"
	"github.com/eoussama/anusic-api/src/shared/models"
)

// Logging dump location
var cacheDirectory = "data"
var cacheFile = "cache.json"

// Cache caches all scraped data
var Cache models.Cache = models.Cache{}

// LoadCache loads the cache files
func LoadCache() bool {
	Log("Loading cache...", enums.LogInfo)

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
	Log("Saving cache...", enums.LogInfo)

	// Constructing the cache file
	absPath, _ := filepath.Abs(".")
	directoryPath := filepath.Join(absPath, cacheDirectory)
	filePath := filepath.Join(directoryPath, cacheFile)

	// Marshalling the data
	file, _ := json.MarshalIndent(cache, "", " ")

	// Writing to the export file
	_ = os.Mkdir(directoryPath, 0755)
	_ = ioutil.WriteFile(filePath, file, 0644)

	Log("Cache saved in "+filePath, enums.LogInfo)
}
