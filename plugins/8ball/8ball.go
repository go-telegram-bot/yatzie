package magicball

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"
	"log"
	"strings"
)

var quips = []string{
	"It is certain", "It is decidedly so", "Without a doubt",
	"Yes definitely", "You may rely on it", "As I see it, yes",
	"Most likely", "Outlook good", "Yes", "Signs point to yes",
	"Reply hazy try again", "Ask again later",
	"Better not tell you now", "Cannot predict now",
	"Concentrate and ask again", "Don\"t count on it",
	"My reply is no", "My sources say no", "Outlook not so good",
	"Very doubtful",
}

type MagicBallPlugin struct {
}

func (m *MagicBallPlugin) OnStart() {
	log.Println("[MagicBallPlugin] Started")
}

func (m *MagicBallPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"8ball") {
		bot.SendMessage(message.Chat,
			util.RandomFromArray(quips), nil)
	}
}

func init() {
	plugin_registry.RegisterPlugin(&MagicBallPlugin{})
	plugin_registry.RegisterCommand("8ball", "Ask me a question")
}
