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

package imdb

import (
	"encoding/json"

	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"io"
	"log"
	"strings"

	"gopkg.in/tucnak/telebot.v1"
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

type IMDBPlugin struct {
	//whatever
}

func init() {
	plugin_registry.RegisterPlugin(&IMDBPlugin{})
}

func (m *IMDBPlugin) OnStart() {
	log.Println("[IMDB] Started")
	plugin_registry.RegisterCommand("imdb", "Search a movie on imdb")

}

func (m *IMDBPlugin) OnStop() {
	plugin_registry.UnregisterCommand("imdb")
}

func (m *IMDBPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"imdb") {
		imdbsearch := message.Text
		log.Println("Searching " + imdbsearch)
		imdbsearch = strings.Replace(imdbsearch, config.CommandPrefix+"imdb ", "", -1)
		imdbsearch = strings.Replace(imdbsearch, " ", "%20", -1)

		util.DecodeJson("http://www.imdbapi.com/?t="+imdbsearch, func(body io.ReadCloser) bool {
			var imdb Movie
			err := json.NewDecoder(body).Decode(&imdb)

			bot.SendMessage(message.Chat, imdb.Title+" - "+imdb.Year+"\n"+imdb.Genre+" - "+imdb.Duration+"\n"+imdb.Image+"\n"+"http://imdb.com/title/"+imdb.Id, nil)

			if err != nil {
				return false
			} else {
				return true
			}
		})

	}

}
