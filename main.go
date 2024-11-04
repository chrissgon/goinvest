package main

import (
	"log"

	"github.com/chrissgon/goinvest/port/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot.StartBot()
}
