// Copyright 2017-2018 Ettore Di Giacinto
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package help

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"gopkg.in/tucnak/telebot.v1"

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
