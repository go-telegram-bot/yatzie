package dogr

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"strings"
)

type DogrPlugin struct{}

func (m *DogrPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
	if strings.Contains(message.Text, config.CommandPrefix+"doge") {
		doge := message.Text
		doge = strings.Replace(doge, config.CommandPrefix+"doge ", "", -1)
		doge = strings.Replace(doge, " ", "/", -1)
		//bot.SendMessage(message.Chat,
		//	"http://dogr.io/"+doge+".png", nil)
		util.SendPhoto("http://dogr.io/"+doge+".png", message, bot)
	}
}

func init() {
	plugin_registry.RegisterPlugin(&DogrPlugin{})
	plugin_registry.RegisterCommand("doge bla bla bla", "Generate a doge with your text")
}
