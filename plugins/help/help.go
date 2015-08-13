package help

import (
	"bytes"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/tucnak/telebot"
	"math/rand"
	"sort"
)

var quips = []string{
	"Hello stranger",
	"Hello Master",
	"Fuck you bitch",
	"I COMMAND YOU, REMEMBER THAT",
}

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config plugin_registry.Config, message telebot.Message) {
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
			buffer.WriteString(config.CommandPrefix + v + " - " + plugin_registry.Commands[v] + "\n")
		}

		bot.SendMessage(message.Chat,
			RandomQuip()+" "+message.Sender.FirstName+"\n Those are my commands: \n"+buffer.String(), nil)
	}
}
func RandomQuip() string {
	return quips[rand.Intn(len(quips))]
}
func init() {
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("help", "Display this help")
}
