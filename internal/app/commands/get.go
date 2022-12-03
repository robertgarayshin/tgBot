package commands

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *botAPI.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	prod, err := c.productService.Get(idx - 1)
	if err != nil {
		log.Println("can't get the product", idx, err)
		return
	}

	msg := botAPI.NewMessage(inputMessage.Chat.ID, prod.Title)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
