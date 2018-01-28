// Copyright 2017-2018 Ettore Di Giacinto
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package yatziebot

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"gopkg.in/tucnak/telebot.v1"

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
