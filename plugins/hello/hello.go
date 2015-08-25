package hello

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/tucnak/telebot"

	"log"
)

type HelloPlugin struct {
}

func init() {
	plugin_registry.RegisterPlugin(&HelloPlugin{})
}

func (m *HelloPlugin) OnStart() {
	log.Println("[HelloPlugin] Started")
	plugin_registry.RegisterCommand("hi", "Says hello")

}

func (m *HelloPlugin) OnStop() {
	plugin_registry.UnregisterCommand("hi")
}

func (m *HelloPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if message.Text == config.CommandPrefix+"hi" {
		bot.SendMessage(message.Chat,
			"Hello, "+message.Sender.FirstName+"!", nil)
	}
}
