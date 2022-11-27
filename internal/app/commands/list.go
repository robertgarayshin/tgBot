package commands

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

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
