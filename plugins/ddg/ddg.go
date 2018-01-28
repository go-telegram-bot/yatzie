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

package ddg

import (
	"fmt"

	"github.com/ajanicij/goduckgo/goduckgo"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"gopkg.in/tucnak/telebot.v1"

	"log"
	"strings"
)

type DDGPlugin struct{}

func init() {
	plugin_registry.RegisterPlugin(&DDGPlugin{})
}

func (m *DDGPlugin) OnStart() {
	log.Println("[DDGPlugin] Started")

	plugin_registry.RegisterCommand("ddg", "Search in DuckDuckGo")

}

func (m *DDGPlugin) OnStop() {
	log.Println("[DDGPlugin] Stopped")
	plugin_registry.UnregisterCommand("ddg")
}

func (m *DDGPlugin) Run(bot *telebot.Bot, message telebot.Message) {
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"ddg") {
		args := util.StripPluginCommand(message.Text, config.CommandPrefix, "ddg")
		bot.SendMessage(message.Chat, SearchCmd(args), nil)

	}

}

// WebSearch takes a query string as an argument and returns
// a formatted string containing the results from DuckDuckGo.
func SearchCmd(query string) string {
	msg, err := goduckgo.Query(query)
	if err != nil {
		return fmt.Sprintf("DDG Error: %v\n", err)
	}

	switch {
	case len(msg.RelatedTopics) > 0:
		return fmt.Sprintf("First Topical Result: [ %s ]( %s )\n", msg.RelatedTopics[0].FirstURL, msg.RelatedTopics[0].Text)
	case len(msg.Results) > 0:
		return fmt.Sprintf("First External result: [ %s ]( %s )\n", msg.Results[0].FirstURL, msg.Results[0].Text)
	case len(msg.Redirect) > 0:
		return fmt.Sprintf("Redirect result: %s\n", msg.Redirect)
	default:
		return fmt.Sprintf("Query: '%s' returned no results.\n", query)
	}
}
