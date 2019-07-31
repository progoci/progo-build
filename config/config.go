package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Init loads a .env file in root directory.
func Init() {
	godotenv.Load()
}

// Get retrieves a environment variable.
func Get(key string) string {

	return os.Getenv(key)
}
