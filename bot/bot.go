package yatziebot

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	"log"
	"strconv"
)

type YatzieBot struct {
	Bot *telebot.Bot
}

func NewBot(config util.Config) (*YatzieBot, error) {
	bot, err := telebot.NewBot(config.Token)

	if err != nil {
		return nil, err
	}

	plugin_registry.Config = config
	plugin_registry.Bot = bot

	// Bootstrapper for plugins
	for _, d := range plugin_registry.Plugins {
		go d.OnStart()
	}
	log.Println(strconv.Itoa(len(plugin_registry.Plugins)) + " plugins loaded")

	bot.Messages = make(chan telebot.Message, 100)
	bot.Queries = make(chan telebot.Query, 1000)

	go messages(bot)
	go queries(bot)
	return &YatzieBot{Bot: bot}, nil
}

func messages(bot *telebot.Bot) {
	for message := range bot.Messages {
		for _, d := range plugin_registry.Plugins {

			if obj, ok := d.(interface {
				Run(*telebot.Bot, telebot.Message)
			}); ok {
				go obj.Run(bot, message)
			}

			if obj, ok := d.(interface {
				Run(telebot.Message)
			}); ok {
				go obj.Run(message)
			}

		}
	}
}

func queries(bot *telebot.Bot) {
	for query := range bot.Queries {
		for _, d := range plugin_registry.Plugins {

			if obj, ok := d.(interface {
				Query(*telebot.Bot, telebot.Query)
			}); ok {
				go obj.Query(bot, query)
			}

			if obj, ok := d.(interface {
				Query(telebot.Query)
			}); ok {
				go obj.Query(query)
			}
		}
	}
}
