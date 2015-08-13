package echo

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/tucnak/telebot"
	"log"
	"strconv"
)

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config plugin_registry.Config, message telebot.Message) {
	log.Println(">> ID: [" + strconv.Itoa(message.Sender.ID) + " ] Name: [" + message.Sender.FirstName + " " + message.Sender.LastName + "] Username: [" + message.Sender.Username + "]\n\tsaid: " + message.Text)
}

func init() {
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
}
