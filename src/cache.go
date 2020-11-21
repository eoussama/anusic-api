package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var cachedAnimeList []Anime = []Anime{}

func loadCache() bool {
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
	json.Unmarshal(byteValue, &cachedAnimeList)

	// Clearing resources
	defer jsonFile.Close()
	return true
}
