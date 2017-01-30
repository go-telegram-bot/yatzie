package main

// Plugins gets automatically loaded on import
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot/yatzie/bot"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	. "github.com/mattn/go-getopt"

	_ "github.com/go-telegram-bot/yatzie/plugins/8ball"
	_ "github.com/go-telegram-bot/yatzie/plugins/admin"
	_ "github.com/go-telegram-bot/yatzie/plugins/ddg"
	_ "github.com/go-telegram-bot/yatzie/plugins/dogr"
	_ "github.com/go-telegram-bot/yatzie/plugins/echo"
	_ "github.com/go-telegram-bot/yatzie/plugins/hal"
	_ "github.com/go-telegram-bot/yatzie/plugins/hello"
	_ "github.com/go-telegram-bot/yatzie/plugins/help"
	_ "github.com/go-telegram-bot/yatzie/plugins/imdb"
	_ "github.com/go-telegram-bot/yatzie/plugins/norris"
	_ "github.com/go-telegram-bot/yatzie/plugins/nsfw"
	_ "github.com/go-telegram-bot/yatzie/plugins/xkcd"
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
	fmt.Println("Configuration file: " + configurationFile)
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

	if config.Token != "" {
		fmt.Println("Token: " + config.Token)
	}

	if logFile != "" {
		fmt.Println("Log file: " + logFile)
	}

	YatzieBot, err := yatziebot.NewBot(config)
	if err != nil {
		log.Fatal("error spawning bot: %v", err)
	}

	YatzieBot.Bot.Start(1 * time.Second)

}
