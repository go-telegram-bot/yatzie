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

package dogr

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"gopkg.in/tucnak/telebot.v1"

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
