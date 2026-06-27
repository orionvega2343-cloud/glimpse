package main

import (
	"glimpse/internal/overlay"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//Загрузка .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//UX окно управления
	overlay.Overlay()

}
