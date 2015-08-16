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

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
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

}

func init() {
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("imgsearch", "Search images on google")
	plugin_registry.RegisterCommand("search", "Search on google")

}
