package utils

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnvVars loads the environment variables
func LoadEnvVars() {
	log.Println("Loading env vars...")

	absPath, _ := filepath.Abs(".")
	path := filepath.Join(absPath, "config", ".env")

	err := godotenv.Load(path)
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
