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

package hello

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"gopkg.in/tucnak/telebot.v1"

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
