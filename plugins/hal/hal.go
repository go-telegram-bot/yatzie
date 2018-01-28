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

// The hal plugin uses microhal to give the bot a little bit of soul
// If mentioned or talked directly, this plugin make your bot answer
package hal

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	cobe "github.com/mudler/go.cobe"
	"gopkg.in/tucnak/telebot.v1"

	"log"
	"strings"
)

type HalPlugin struct {
	Brain *cobe.Cobe2Brain
}

func init() {
	plugin_registry.RegisterPlugin(&HalPlugin{})
}

func (m *HalPlugin) OnStart() {
	b, err := cobe.OpenCobe2Brain(plugin_registry.Config.BrainFile)
	m.Brain = b
	if err != nil {
		log.Fatalf("Opening brain file: %s", err)
	}
	log.Println("[HalPlugin] Started")
}

func (m *HalPlugin) OnStop() {
	log.Println("[HalPlugin] Disabled")
}

func (m *HalPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config

	if !strings.HasPrefix(message.Text, config.CommandPrefix) && !util.MatchAnyURL(message.Text) {
		text := strings.Replace(message.Text, "@"+bot.Identity.Username, "", -1)
		log.Println("[HalPlugin] Learning " + text)
		m.Brain.Learn(text)
		if message.IsPersonal() || strings.Contains(message.Text, "@"+bot.Identity.Username) {
			bot.SendMessage(message.Chat, m.Brain.Reply(text), nil)
		}

	}
}
