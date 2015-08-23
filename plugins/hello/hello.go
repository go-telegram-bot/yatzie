package hello

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/tucnak/telebot"

	"log"
)

type HelloPlugin struct {
}

func (m *HelloPlugin) OnStart() {
	log.Println("[HelloPlugin] Started")
}

func (m *HelloPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if message.Text == config.CommandPrefix+"hi" {
		bot.SendMessage(message.Chat,
			"Hello, "+message.Sender.FirstName+"!", nil)
	}
}

func init() {
	plugin_registry.RegisterPlugin(&HelloPlugin{})
	plugin_registry.RegisterCommand("hi", "Says hello")
}
