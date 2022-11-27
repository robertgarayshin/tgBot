package commands

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robertgarayshin/tgBot/internal/service/product"
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
