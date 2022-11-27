package commands

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robertgarayshin/tgBot/internal/service/product"
	"log"
)

type Commander struct {
	bot            *botAPI.BotAPI
	productService *product.Service
}

func NewCommander(bot *botAPI.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) Help(inputMessage *botAPI.Message) {
	msg := botAPI.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func (c *Commander) List(inputMessage *botAPI.Message) {

	outputMsgText := "Here all the products: \n\n"

	products := c.productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	msg := botAPI.NewMessage(inputMessage.Chat.ID, outputMsgText)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func (c *Commander) Default(inputMessage *botAPI.Message) {
	msg := botAPI.NewMessage(inputMessage.Chat.ID, "Written "+inputMessage.Text)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
