package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	. "github.com/mattn/go-getopt"
	"github.com/tucnak/telebot"

	_ "github.com/go-telegram-bot/yatzie/plugins/dogr"
	_ "github.com/go-telegram-bot/yatzie/plugins/google"

	_ "github.com/go-telegram-bot/yatzie/plugins/8ball"
	_ "github.com/go-telegram-bot/yatzie/plugins/echo"
	_ "github.com/go-telegram-bot/yatzie/plugins/hal"
	_ "github.com/go-telegram-bot/yatzie/plugins/hello"
	_ "github.com/go-telegram-bot/yatzie/plugins/help"
	_ "github.com/go-telegram-bot/yatzie/plugins/xkcd"

	_ "github.com/go-telegram-bot/yatzie/plugins/imdb"
	_ "github.com/go-telegram-bot/yatzie/plugins/norris"
	_ "github.com/go-telegram-bot/yatzie/plugins/nsfw"
)

func main() {

	var c int
	var configurationFile = "telegram-config.json"
	var logFile string
	OptErr = 0
	for {
		if c = Getopt("c:l:h"); c == EOF {
			break
		}
		switch c {
		case 'c':
			configurationFile = OptArg
		case 'l':
			logFile = OptArg
		case 'h':
			println("usage: " + os.Args[0] + " [-c configfile.json|-l logfile|-h]")
			os.Exit(1)
		}
	}

	config, err := util.LoadConfig(configurationFile)

	if logFile != "" {
		//Set logging to file
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	bot, err := telebot.NewBot(config.Token)
	if config.Token != "" {
		fmt.Println("Token: " + config.Token)
	}
	fmt.Println("Configuration file: " + configurationFile)
	fmt.Println("Log file: " + logFile)

	if err != nil {
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)
	plugin_registry.Config = config
	plugin_registry.Bot = bot

	for message := range messages {
		for _, d := range plugin_registry.Plugins {
			go d.Run(message)
		}
	}
}
