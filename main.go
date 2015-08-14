package main

import (
	"fmt"
	"time"

	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	_ "github.com/go-telegram-bot/yatzie/plugins/9gag"
	_ "github.com/go-telegram-bot/yatzie/plugins/dogr"
	_ "github.com/go-telegram-bot/yatzie/plugins/echo"
	_ "github.com/go-telegram-bot/yatzie/plugins/hal"
	_ "github.com/go-telegram-bot/yatzie/plugins/hello"
	_ "github.com/go-telegram-bot/yatzie/plugins/help"
	_ "github.com/go-telegram-bot/yatzie/plugins/norris"
	_ "github.com/go-telegram-bot/yatzie/plugins/nsfw"
)

func main() {
	config, err := util.LoadConfig("telegram-config.json")
	bot, err := telebot.NewBot(config.Token)
	fmt.Println("Token: " + config.Token)
	if err != nil {
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		for _, d := range plugin_registry.Plugins {
			go d.Run(bot, config, message)
		}
	}
}
