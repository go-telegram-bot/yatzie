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

package nsfw

import (
	"encoding/json"

	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"io"
	"log"
	"math/rand"
	"strings"

	"gopkg.in/tucnak/telebot.v1"
)

type ImageJson struct {
	Id      int    `json:"id"`
	Preview string `json:"preview"`
}
type ImagesJson []ImageJson

type HentaiImage struct {
	Id  int    `json:"id"`
	Url string `json:"file_url"`
}
type HentaiJson []HentaiImage

type NSFWPlugin struct {
}

func init() {
	plugin_registry.RegisterPlugin(&NSFWPlugin{})

}

func (m *NSFWPlugin) OnStart() {
	log.Println("[NSFWPlugin] Started")
	plugin_registry.RegisterCommand("hentai", "Display some random hentai image")
	plugin_registry.RegisterCommand("boobs", "Display some random BOOTY image")
	plugin_registry.RegisterCommand("butts", "Display some random BUUTTTSY image")
}

func (m *NSFWPlugin) OnStop() {
	plugin_registry.UnregisterCommand("hentai")
	plugin_registry.UnregisterCommand("boobs")
	plugin_registry.UnregisterCommand("butts")
}

func (m *NSFWPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if message.Text == config.CommandPrefix+"butts" {

		SendNSFWButt(message, bot)

	}

	if message.Text == config.CommandPrefix+"hentai" {

		SendNSFWHentai(message, bot)
		//boobs, _ := getHentai("http://danbooru.donmai.us/posts.json")
		//boobs2, _ := getHentai("http://danbooru.donmai.us/explore/posts/popular.json")

		/*
			unionboobs := make(HentaiJson, len(boobs)+len(boobs2))
			copy(unionboobs, boobs)
			copy(unionboobs[len(boobs):], boobs2)
			url := "http://danbooru.donmai.us"
			util.SendPhoto(url+unionboobs[rand.Intn(len(unionboobs))].Url, message, bot)
		*/
		//bot.SendMessage(message.Chat,
		//	url+unionboobs[rand.Intn(len(unionboobs))].Url, nil)
	}

	if message.Text == config.CommandPrefix+"boobs" {

		SendNSFWBoob(message, bot)

	}
}

func SendNSFWBoob(message telebot.Message, bot *telebot.Bot) {
	util.DecodeJson("http://api.oboobs.ru/noise/1", func(body io.ReadCloser) bool {
		var data ImagesJson
		err := json.NewDecoder(body).Decode(&data)
		url := "http://media.oboobs.ru/"

		for _, i := range data {
			//bot.SendMessage(message.Chat, url+strings.Replace(i.Preview, "_preview", "", -1), nil)
			util.SendPhoto(url+strings.Replace(i.Preview, "_preview", "", -1), message, bot)

		}

		if err != nil {
			return false
		} else {
			return true
		}
	})
}

func SendNSFWButt(message telebot.Message, bot *telebot.Bot) {
	util.DecodeJson("http://api.obutts.ru/noise/1", func(body io.ReadCloser) bool {
		var data ImagesJson
		err := json.NewDecoder(body).Decode(&data)
		url := "http://media.obutts.ru/"

		for _, i := range data {
			//bot.SendMessage(message.Chat, url+strings.Replace(i.Preview, "_preview", "", -1), nil)
			util.SendPhoto(url+strings.Replace(i.Preview, "_preview", "", -1), message, bot)

		}

		if err != nil {
			return false
		}
		return true

	})
}

func SendNSFWHentai(message telebot.Message, bot *telebot.Bot) {

	util.DecodeJson("http://danbooru.donmai.us/posts.json", func(body io.ReadCloser) bool {
		var data HentaiJson
		err := json.NewDecoder(body).Decode(&data)
		url := "http://danbooru.donmai.us"

		util.SendPhoto(url+data[rand.Intn(len(data))].Url, message, bot)

		if err != nil {
			return false
		}
		return true

	})

	util.DecodeJson("http://danbooru.donmai.us/explore/posts/popular.json", func(body io.ReadCloser) bool {
		var data HentaiJson
		err := json.NewDecoder(body).Decode(&data)
		url := "http://danbooru.donmai.us"

		util.SendPhoto(url+data[rand.Intn(len(data))].Url, message, bot)

		if err != nil {
			return false
		}
		return true

	})
}
