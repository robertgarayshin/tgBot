package main

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/robertgarayshin/tgBot/internal/app/commands"
	"github.com/robertgarayshin/tgBot/internal/service/product"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("token")
	bot, err := botAPI.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := botAPI.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)
	for update := range updates {
		if update.Message == nil {
			continue
		} // If we got not a message
		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.Default(update.Message)
		}

	}
}
