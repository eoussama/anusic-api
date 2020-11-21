package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var cachedAnimeList []Anime = []Anime{}

func loadCache() bool {

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

	// Sanitizing alt name
	for index := 0; index < len(cachedAnimeList); index++ {
		cachedAnimeList[index].Name = strings.Replace(cachedAnimeList[index].Name, "\"", "", -1)
		cachedAnimeList[index].AltName = strings.Replace(cachedAnimeList[index].AltName, "\"", "", -1)
	}

	// Clearing resources
	defer jsonFile.Close()
	return true
}
