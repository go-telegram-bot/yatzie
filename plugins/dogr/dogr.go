package dogr

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"strings"
)

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config plugin_registry.Config, message telebot.Message) {
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
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("doge bla bla bla", "Generate a doge with your text")

}
