package dogr

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"

	"log"
	"strings"
)

type DogrPlugin struct{}

func init() {
	plugin_registry.RegisterPlugin(&DogrPlugin{})
}

func (m *DogrPlugin) OnStart() {
	log.Println("[DogrPlugin] Started")
	plugin_registry.RegisterCommand("doge bla bla bla", "Generate a doge with your text")

}

func (m *DogrPlugin) OnStop() {
	log.Println("[DogrPlugin] Stopped")
	plugin_registry.UnregisterCommand("doge")

}

func (m *DogrPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"doge") {
		doge := message.Text
		doge = strings.Replace(doge, config.CommandPrefix+"doge ", "", -1)
		doge = strings.Replace(doge, " ", "/", -1)
		//bot.SendMessage(message.Chat,
		//	"http://dogr.io/"+doge+".png", nil)
		util.SendPhoto("http://dogr.io/"+doge+".png", message, bot)
	}
}
