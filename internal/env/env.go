package env

import (
	"log"

	"github.com/lpernett/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return err
}
