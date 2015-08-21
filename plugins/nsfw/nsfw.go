package nsfw

import (
	"encoding/json"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"io"
	"math/rand"
	"strings"
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

func (m *NSFWPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {

	if message.Text == config.CommandPrefix+"butts" {

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
			} else {
				return true
			}
		})

	}

	if message.Text == config.CommandPrefix+"hentai" {
		//boobs, _ := getHentai("http://danbooru.donmai.us/posts.json")
		//boobs2, _ := getHentai("http://danbooru.donmai.us/explore/posts/popular.json")

		util.DecodeJson("http://danbooru.donmai.us/posts.json", func(body io.ReadCloser) bool {
			var data HentaiJson
			err := json.NewDecoder(body).Decode(&data)
			url := "http://danbooru.donmai.us"

			util.SendPhoto(url+data[rand.Intn(len(data))].Url, message, bot)

			if err != nil {
				return false
			} else {
				return true
			}
		})

		util.DecodeJson("http://danbooru.donmai.us/explore/posts/popular.json", func(body io.ReadCloser) bool {
			var data HentaiJson
			err := json.NewDecoder(body).Decode(&data)
			url := "http://danbooru.donmai.us"

			util.SendPhoto(url+data[rand.Intn(len(data))].Url, message, bot)

			if err != nil {
				return false
			} else {
				return true
			}
		})
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
}

func init() {
	plugin_registry.RegisterPlugin(&NSFWPlugin{})
	plugin_registry.RegisterCommand("hentai", "Display some random hentai image")
	plugin_registry.RegisterCommand("boobs", "Display some random BOOTY image")
	plugin_registry.RegisterCommand("butts", "Display some random BUUTTTSY image")
}
