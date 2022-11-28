package commands

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Get(inputMessage *botAPI.Message) {
	msg := botAPI.NewMessage(inputMessage.Chat.ID, "TBD")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
