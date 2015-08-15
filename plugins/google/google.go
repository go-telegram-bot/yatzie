package google

import (
	"encoding/json"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"log"
	"net/http"
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

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
	if strings.Contains(message.Text, config.CommandPrefix+"imgsearch") {
		imgsearch := message.Text
		log.Println("Searching " + imgsearch)
		imgsearch = strings.Replace(imgsearch, config.CommandPrefix+"imgsearch ", "", -1)
		imgsearch = strings.Replace(imgsearch, " ", "%20", -1)

		imgs, _ := search("https://ajax.googleapis.com/ajax/services/search/images?v=1.0&q=" + imgsearch)

		for _, i := range imgs.Response.Results {
			util.SendPhoto(i.Url, message, bot)
			log.Println("Found " + i.Title)
		}

	}

	if strings.Contains(message.Text, config.CommandPrefix+"search") {
		websearch := message.Text
		log.Println("Searching " + websearch)
		websearch = strings.Replace(websearch, config.CommandPrefix+"search ", "", -1)
		websearch = strings.Replace(websearch, " ", "%20", -1)

		urls, _ := search("https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=" + websearch)

		for _, i := range urls.Response.Results {
			bot.SendMessage(message.Chat, i.Title+" - "+i.Url, nil)
			log.Println("Found " + i.Title)
		}

	}

}

func search(url string) (Response, error) {
	var data Response
	r, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	return data, err
}

func init() {
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("imgsearch", "Search images on google")
	plugin_registry.RegisterCommand("search", "Search on google")

}
