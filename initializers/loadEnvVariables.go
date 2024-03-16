package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load enviroment variables from .env file
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
