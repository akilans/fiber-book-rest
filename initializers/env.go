package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load ENV variables from .env file
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Loaded env successfully")
	}
}
