package commands

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(inputMessage *botAPI.Message) {
	msg := botAPI.NewMessage(inputMessage.Chat.ID, "Written "+inputMessage.Text)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func (c *Commander) HandleUpdate(update botAPI.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.Message == nil {
		return
	} // If we got not a message
	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
