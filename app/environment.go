package app

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while loading env file %v", err)
		panic(err)
	}
}
