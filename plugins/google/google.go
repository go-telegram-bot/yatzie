package google

import (
	"encoding/json"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"io"
	"log"
	"strings"
)

type Response struct {
	Response Results `json:"responseData"`
}

type Results struct {
	Results []Result `json:"results"`
}

type Result struct {
	Title string `json:"titleNoFormatting"`
	Url   string `json:"url"`
}

type GooglePlugin struct {
}

func (m *GooglePlugin) OnStart() {
	log.Println("[GooglePlugin] Started")
}

func (m *GooglePlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"imgsearch") {
		imgsearch := message.Text
		log.Println("Searching " + imgsearch)
		imgsearch = strings.Replace(imgsearch, config.CommandPrefix+"imgsearch ", "", -1)
		imgsearch = strings.Replace(imgsearch, " ", "%20", -1)

		util.DecodeJson("https://ajax.googleapis.com/ajax/services/search/images?v=1.0&q="+imgsearch, func(body io.ReadCloser) bool {
			var data Response
			err := json.NewDecoder(body).Decode(&data)

			for _, i := range data.Response.Results {
				util.SendPhoto(i.Url, message, bot)
				log.Println("Found " + i.Title)
			}

			if err != nil {
				return false
			} else {
				return true
			}
		})

	}

	if strings.Contains(message.Text, config.CommandPrefix+"search") {
		websearch := message.Text
		log.Println("Searching " + websearch)
		websearch = strings.Replace(websearch, config.CommandPrefix+"search ", "", -1)
		websearch = strings.Replace(websearch, " ", "%20", -1)

		util.DecodeJson("https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q="+websearch, func(body io.ReadCloser) bool {
			var data Response
			err := json.NewDecoder(body).Decode(&data)

			for _, i := range data.Response.Results {
				bot.SendMessage(message.Chat, i.Title+" - "+i.Url, nil)
				log.Println("Found " + i.Title)
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
	plugin_registry.RegisterPlugin(&GooglePlugin{})
	plugin_registry.RegisterCommand("imgsearch", "Search images on google")
	plugin_registry.RegisterCommand("search", "Search on google")
}
