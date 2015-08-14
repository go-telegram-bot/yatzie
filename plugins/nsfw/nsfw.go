package nsfw

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"log"
	"math/rand"
	"net/http"
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

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {

	if message.Text == config.CommandPrefix+"butts" {
		boobs, _ := getImages("http://api.obutts.ru/noise/1")
		url := "http://media.obutts.ru/"

		for _, i := range boobs {
			//bot.SendMessage(message.Chat, url+strings.Replace(i.Preview, "_preview", "", -1), nil)
			util.SendPhoto(url+strings.Replace(i.Preview, "_preview", "", -1), message, bot)

		}

	}

	if message.Text == config.CommandPrefix+"hentai" {
		boobs, _ := getHentai("http://danbooru.donmai.us/posts.json")
		boobs2, _ := getHentai("http://danbooru.donmai.us/explore/posts/popular.json")
		unionboobs := make(HentaiJson, len(boobs)+len(boobs2))
		copy(unionboobs, boobs)
		copy(unionboobs[len(boobs):], boobs2)
		url := "http://danbooru.donmai.us"
		util.SendPhoto(url+unionboobs[rand.Intn(len(unionboobs))].Url, message, bot)
		//bot.SendMessage(message.Chat,
		//	url+unionboobs[rand.Intn(len(unionboobs))].Url, nil)
	}

	if message.Text == config.CommandPrefix+"boobs" {
		boobs, _ := getImages("http://api.oboobs.ru/noise/1")
		url := "http://media.oboobs.ru/"

		for _, i := range boobs {
			//bot.SendMessage(message.Chat, url+strings.Replace(i.Preview, "_preview", "", -1), nil)
			util.SendPhoto(url+strings.Replace(i.Preview, "_preview", "", -1), message, bot)
		}

	}
}

func getHentai(url string) (HentaiJson, error) {
	var data HentaiJson
	r, err := http.Get(url)
	fmt.Println(url)

	if err != nil {
		return data, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	for _, i := range data {
		log.Println(
			url + i.Url)
	}
	return data, err
}
func getImages(url string) (ImagesJson, error) {
	var data ImagesJson
	r, err := http.Get(url)
	fmt.Println(url)

	if err != nil {
		return data, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	for _, i := range data {
		log.Println(
			url + i.Preview)
	}
	return data, err
}

func init() {
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("hentai", "Display some random hentai image")
	plugin_registry.RegisterCommand("boobs", "Display some random BOOTY image")
	plugin_registry.RegisterCommand("butts", "Display some random BUUTTTSY image")

}
