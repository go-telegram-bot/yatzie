package hello

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"
)

type HelloPlugin struct {
}

func (m *HelloPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
	if message.Text == config.CommandPrefix+"hi" {
		bot.SendMessage(message.Chat,
			"Hello, "+message.Sender.FirstName+"!", nil)
	}
}

func init() {
	plugin_registry.RegisterPlugin(&HelloPlugin{})
	plugin_registry.RegisterCommand("hi", "Says hello")
}
