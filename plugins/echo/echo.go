package echo

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/tucnak/telebot"
	"log"
	"strconv"
)

type EchoPlugin struct {
}

func (m *EchoPlugin) OnStart() {
	log.Println("[EchoPlugin] Started")
}

func (m *EchoPlugin) Run(message telebot.Message) {
	log.Println(">> ID: [" + strconv.Itoa(message.Sender.ID) + " ] Name: [" + message.Sender.FirstName + " " + message.Sender.LastName + "] Username: [" + message.Sender.Username + "]\n\tsaid: " + message.Text)
}

func init() {
	plugin_registry.RegisterPlugin(&EchoPlugin{})
}
