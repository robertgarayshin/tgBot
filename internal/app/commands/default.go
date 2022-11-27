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
