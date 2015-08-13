package hello

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"

	"github.com/tucnak/telebot"
)

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config plugin_registry.Config, message telebot.Message) {
	if message.Text == config.CommandPrefix+"hi" {
		bot.SendMessage(message.Chat,
			"Hello, "+message.Sender.FirstName+"!", nil)
	}
}

func init() {
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("hi", "Says hello")

}
