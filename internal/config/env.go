package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {

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
			break
		}
		dir = parent
	}

	return ""
}
