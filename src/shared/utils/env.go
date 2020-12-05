package utils

import (
	"path/filepath"

	"github.com/eoussama/anusic-api/src/shared/enums"
	"github.com/joho/godotenv"
)

// LoadEnvVars loads the environment variables
func LoadEnvVars() {
	Log("Loading env vars...", enums.LogInfo)

	absPath, _ := filepath.Abs(".")
	path := filepath.Join(absPath, "config", ".env")

	err := godotenv.Load(path)
	if err != nil {
		Log("Error loading .env file", enums.LogError)
	}
}
