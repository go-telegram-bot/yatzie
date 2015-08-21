package norris

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/tucnak/telebot"
	"net/http"
)

type Response struct {
	Type  int  `json:"type"`
	Value Joke `json:"value"`
}

type Joke struct {
	Id         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

type NorrisPlugin struct {
}

func (m *NorrisPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if message.Text == config.CommandPrefix+"norris" {
		res, _ := getJoke("http://api.icndb.com/jokes/random")
		bot.SendMessage(message.Chat, res.Value.Joke, nil)

	}

}

func getJoke(url string) (Response, error) {
	var data Response
	r, err := http.Get(url)
	fmt.Println(url)

	if err != nil {
		return data, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	return data, err
}

func init() {
	plugin_registry.RegisterPlugin(&NorrisPlugin{})
	plugin_registry.RegisterCommand("norris", "Get a kicking ass chuck norris quote!")
}
