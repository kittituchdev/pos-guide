package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnv(envName string) string {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(filepath.Join(".", ".env"))
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv(envName)
}
