package imdb

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Duration string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Actors   string `json:"Actors"`
	Id       string `json:"imdbID"`
	Image    string `json:"Poster"`
}

type MyPlugin struct {
	//whatever
}

func (m *MyPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
	if strings.Contains(message.Text, config.CommandPrefix+"imdb") {
		imdbsearch := message.Text
		log.Println("Searching " + imdbsearch)
		imdbsearch = strings.Replace(imdbsearch, config.CommandPrefix+"imdb ", "", -1)
		imdbsearch = strings.Replace(imdbsearch, " ", "%20", -1)

		imdb, _ := search("http://www.imdbapi.com/?t=" + imdbsearch)

		bot.SendMessage(message.Chat, imdb.Title+" - "+imdb.Year+"\n"+imdb.Genre+" - "+imdb.Duration+"\n"+imdb.Image+"\n"+"http://imdb.com/title/"+imdb.Id, nil)

	}

}

func search(url string) (Movie, error) {
	var data Movie
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
	my := &MyPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("imdb", "Search a movie on imdb")

}
