package help

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	"bytes"
	"log"
	"sort"
)

var quips = []string{
	"Hello ",
	"Retards.. ehm REGARDS ",
	"Fuck you",
	"❓",
}

var answ = []string{
	"👺",
	"🙈",
}

type HelpPlugin struct {
}

func init() {
	plugin_registry.RegisterPlugin(&HelpPlugin{})
}

func (m *HelpPlugin) OnStart() {
	log.Println("[HelpPlugin] Started")
	plugin_registry.RegisterCommand("help", "Display this help")

}

func (m *HelpPlugin) OnStop() {
	plugin_registry.UnregisterCommand("help")

}

func (m *HelpPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if message.Text == config.CommandPrefix+"help" || message.Text == config.CommandPrefix+"start" {
		mk := make([]string, len(plugin_registry.Commands))
		i := 0
		for k, _ := range plugin_registry.Commands {
			mk[i] = k
			i++
		}
		sort.Strings(mk)
		var buffer bytes.Buffer

		for _, v := range mk {
			buffer.WriteString("▪️" + config.CommandPrefix + v + " - " + plugin_registry.Commands[v] + "\n")
		}

		bot.SendMessage(message.Chat,
			util.RandomFromArray(quips)+", "+message.Sender.FirstName+" "+util.RandomFromArray(answ)+"\n Those are my commands: \n"+buffer.String(), nil)
	}
}
