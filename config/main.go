package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {

	err := godotenv.Load("./config/environments/local.env")
	if err != nil {
		log.Fatal("Error loading .env file!", err)
	}
	return os.Getenv(key)
}
