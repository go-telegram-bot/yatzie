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

package reddit

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/jzelinskie/geddit"

	"fmt"
	"log"
	"strings"

	"gopkg.in/tucnak/telebot.v1"
)

type RedditFeedPlugin struct {
	Session *geddit.LoginSession
}

func init() {
	plugin_registry.RegisterPlugin(&RedditFeedPlugin{})
}

func (m *RedditFeedPlugin) OnStart() {
	log.Println("[RedditFeedPlugin] Started")
	plugin_registry.RegisterCommand("subreddit", "Show latest submissions")
	// Login to reddit
	if plugin_registry.Config.RedditUser != "" && plugin_registry.Config.RedditPassword != "" {
		m.Session, _ = geddit.NewLoginSession(
			plugin_registry.Config.RedditUser,
			plugin_registry.Config.RedditPassword,
			"gedditAgent v1",
		)

	}

}

func (m *RedditFeedPlugin) OnStop() {
	log.Println("[RedditFeedPlugin] Stopped")
	plugin_registry.UnregisterCommand("subreddit")
}

func (m *RedditFeedPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"subreddit") {
		args := util.StripPluginCommand(message.Text, config.CommandPrefix, "subreddit")

		// Set listing options
		subOpts := geddit.ListingOptions{
			Limit: 5,
		}

		// Get specific subreddit submissions, sorted by new
		submissions, _ := m.Session.SubredditSubmissions(args, geddit.NewSubmissions, subOpts)

		// Print title and author of each submission
		for _, s := range submissions {
			bot.SendMessage(message.Chat, fmt.Sprintf("Title: %s\nThumb:%s\n Url: %s\n Text: %s  \n\n", s.Title, s.ThumbnailURL, s.URL, s.Selftext), nil)
		}

	}
}
