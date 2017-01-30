package reddit

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/jzelinskie/geddit"

	"fmt"
	"github.com/tucnak/telebot"
	"log"
	"strings"
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
