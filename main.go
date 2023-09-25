package main

import (
	"github.com/joho/godotenv"
	"log"
	"rmnvlv/bot-party/bot/actions"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	actions.VkParser("https://api.vk.com/method/wall.get")
}
