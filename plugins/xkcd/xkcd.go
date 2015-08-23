package xkcd

import (
	"encoding/json"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"io"
	"log"
	"strconv"
	"strings"
)

type Result struct {
	Id    int    `json:"num"`
	Title string `json:"safe_title"`
	Img   string `json:"img"`
}

type XkcdPlugin struct {
}

func (m *XkcdPlugin) OnStart() {
	log.Println("[XkcdPlugin] Started")
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

func init() {
	plugin_registry.RegisterPlugin(&XkcdPlugin{})
	plugin_registry.RegisterCommand("xkcd", "Get the latest xkcd")
	plugin_registry.RegisterCommand("xkcd <id>", "show a specific xkcd")
}
