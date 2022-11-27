package main

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
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

	for update := range updates {
		if update.Message == nil {
			continue
		} // If we got not a message
		switch update.Message.Command() {
		case "help":
			printHelp(update.Message, bot)
		case "list":
			printList(update.Message, bot, productService)
		default:
			defaultBehaviour(update.Message, bot)
		}

	}
}

func printHelp(inputMessage *botAPI.Message, bot *botAPI.BotAPI) {
	msg := botAPI.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func printList(inputMessage *botAPI.Message, bot *botAPI.BotAPI, productService *product.Service) {

	outputMsgText := "Here all the products: \n\n"

	products := productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	msg := botAPI.NewMessage(inputMessage.Chat.ID, outputMsgText)
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func defaultBehaviour(inputMessage *botAPI.Message, bot *botAPI.BotAPI) {
	msg := botAPI.NewMessage(inputMessage.Chat.ID, "Written "+inputMessage.Text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
