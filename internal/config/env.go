package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv attempts to load environment variables from .env file if it exists,
// but doesn't fail if the file is not found. This makes it work both in
// development (where .env exists) and in CI/CD (where env vars are set directly).
func LoadEnv() {
	// Try to find the .env file by walking up the directory tree
	// This helps when running tests from different directories
	envPath := findEnvFile()
	if envPath == "" {
		log.Println("No .env file found, using environment variables")
		return
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file from %s: %v", envPath, err)
	}
}

// findEnvFile walks up the directory tree looking for a .env file
func findEnvFile() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root directory
			break
		}
		dir = parent
	}

	return ""
}
