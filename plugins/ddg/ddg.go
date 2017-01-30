package ddg

import (
	"fmt"
	"github.com/ajanicij/goduckgo/goduckgo"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	"log"
	"strings"
)

type DDGPlugin struct{}

func init() {
	plugin_registry.RegisterPlugin(&DDGPlugin{})
}

func (m *DDGPlugin) OnStart() {
	log.Println("[DDGPlugin] Started")

	plugin_registry.RegisterCommand("ddg", "Search in DuckDuckGo")

}

func (m *DDGPlugin) OnStop() {
	log.Println("[DDGPlugin] Stopped")
	plugin_registry.UnregisterCommand("ddg")
}

func (m *DDGPlugin) Run(bot *telebot.Bot, message telebot.Message) {
	config := plugin_registry.Config
	if strings.Contains(message.Text, config.CommandPrefix+"ddg") {
		args := util.StripPluginCommand(message.Text, config.CommandPrefix, "ddg")
		bot.SendMessage(message.Chat, SearchCmd(args), nil)

	}

}

// WebSearch takes a query string as an argument and returns
// a formatted string containing the results from DuckDuckGo.
func SearchCmd(query string) string {
	msg, err := goduckgo.Query(query)
	if err != nil {
		return fmt.Sprintf("DDG Error: %v\n", err)
	}

	switch {
	case len(msg.RelatedTopics) > 0:
		return fmt.Sprintf("First Topical Result: [ %s ]( %s )\n", msg.RelatedTopics[0].FirstURL, msg.RelatedTopics[0].Text)
	case len(msg.Results) > 0:
		return fmt.Sprintf("First External result: [ %s ]( %s )\n", msg.Results[0].FirstURL, msg.Results[0].Text)
	case len(msg.Redirect) > 0:
		return fmt.Sprintf("Redirect result: %s\n", msg.Redirect)
	default:
		return fmt.Sprintf("Query: '%s' returned no results.\n", query)
	}
}
