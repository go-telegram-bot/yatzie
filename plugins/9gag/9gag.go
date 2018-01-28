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

package gag

import (
	"encoding/json"
	"fmt"

	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"log"
	"math/rand"
	"net/http"
	"strings"

	"gopkg.in/tucnak/telebot.v1"
)

type GagJson struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Src   string `json:"src"`
}
type GagsJson []GagJson

type GagPlugin struct {
}

func init() {
	plugin_registry.RegisterPlugin(&GagPlugin{})
}

func (m *GagPlugin) OnStart() {
	log.Println("[GagPlugin] Started")
	plugin_registry.RegisterCommand("gag", "Display some random gag ")
}

func (m *GagPlugin) OnStop() {
	plugin_registry.UnregisterCommand("gag")
}

func (m *GagPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if message.Text == config.CommandPrefix+"gag" {
		gags, err := getImages("http://api-9gag.herokuapp.com/")
		if err != nil {
			//bot.SendMessage(message.Chat, strings.Replace(gags[rand.Intn(len(gags))].Src, `\/`, "/", -1), nil)
			util.SendPhoto(strings.Replace(gags[rand.Intn(len(gags))].Src, `\/`, "/", -1), message, bot)
		} else {
			log.Println("[9gag] error: %v", err)
		}

	}

}

func getImages(url string) (GagsJson, error) {
	var data GagsJson
	r, err := http.Get(url)
	fmt.Println(url)

	if err != nil {
		return data, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	for _, i := range data {
		log.Println(
			url + i.Src)
	}
	return data, err
}
