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

package xkcd

import (
	"encoding/json"

	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"io"
	"log"
	"strconv"
	"strings"

	"gopkg.in/tucnak/telebot.v1"
)

type Result struct {
	Id    int    `json:"num"`
	Title string `json:"safe_title"`
	Img   string `json:"img"`
}

type XkcdPlugin struct {
}

func init() {
	plugin_registry.RegisterPlugin(&XkcdPlugin{})

}

func (m *XkcdPlugin) OnStart() {
	log.Println("[XkcdPlugin] Started")
	plugin_registry.RegisterCommand("xkcd", "Get the latest xkcd")
	plugin_registry.RegisterCommand("xkcd <id>", "show a specific xkcd")
}

func (m *XkcdPlugin) OnStop() {
	plugin_registry.UnregisterCommand("xkcd")
	plugin_registry.UnregisterCommand("xkcd <id>")
}
func (m *XkcdPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"xkcd") {
		xkcd := message.Text
		log.Println("Searching " + xkcd)
		xkcd = strings.Replace(xkcd, config.CommandPrefix+"xkcd ", "", -1)

		util.DecodeJson("https://xkcd.com/"+xkcd+"/info.0.json", func(body io.ReadCloser) bool {
			var data Result
			err := json.NewDecoder(body).Decode(&data)
			bot.SendMessage(message.Chat, strconv.Itoa(data.Id)+" - "+data.Img+" - "+data.Title, nil)
			log.Println("Found " + data.Title)

			if err != nil {
				return false
			} else {
				return true
			}
		})

	}

	if message.Text == config.CommandPrefix+"xkcd" {
		util.DecodeJson("https://xkcd.com/info.0.json", func(body io.ReadCloser) bool {
			var data Result
			err := json.NewDecoder(body).Decode(&data)
			bot.SendMessage(message.Chat, strconv.Itoa(data.Id)+" - "+data.Img+" - "+data.Title, nil)
			log.Println("Found " + data.Title)

			if err != nil {
				return false
			} else {
				return true
			}
		})
	}

}
