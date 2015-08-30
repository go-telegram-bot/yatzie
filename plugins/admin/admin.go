package admin

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"

	"bytes"
	"log"
	"strings"
)

type AdminPlugin struct{}

func init() {
	plugin_registry.RegisterPlugin(&AdminPlugin{})
}

func (m *AdminPlugin) OnStart() {
	log.Println("[AdminPlugin] Started")
}

func (m *AdminPlugin) OnStop() {
	log.Println("[AdminPlugin] Stop")
}

func (m *AdminPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if config.IsAdmin(message.Sender.Username) == false {
		return
	}

	if message.Text == config.CommandPrefix+"help" {
		bot.SendMessage(message.Chat, "👑 👑 👑 👑 HAIL TO THE KING 👑 👑 👑 👑 \n- Admin commands - \n\t"+config.CommandPrefix+"enable <plugin> - Load a specific plugin again in memory\n\t"+config.CommandPrefix+"disable <plugin> - UnLoad a specific plugin in memory\n\t"+config.CommandPrefix+"listplugins - List all plugins\n", nil)

	}

	if message.Text == config.CommandPrefix+"listplugins" {
		ListPlugins(message, bot)

	}
	if strings.Contains(message.Text, config.CommandPrefix+"enable") {
		args := util.StripPluginCommand(message.Text, config.CommandPrefix, "enable")
		if plugin_registry.EnablePlugin(args) {
			bot.SendMessage(message.Chat, args+" Enabled", nil)
		}
		ListPlugins(message, bot)

	}
	if strings.Contains(message.Text, config.CommandPrefix+"disable") {
		args := util.StripPluginCommand(message.Text, config.CommandPrefix, "disable")
		if plugin_registry.DisablePlugin(args) {
			bot.SendMessage(message.Chat, args+" Disabled", nil)

		}
		ListPlugins(message, bot)

	}
}

func ListPlugins(message telebot.Message, bot *telebot.Bot) {
	var loaded bytes.Buffer
	var unloaded bytes.Buffer

	for k, _ := range plugin_registry.Plugins {
		loaded.WriteString("\t" + k + "\n")

	}

	for k, _ := range plugin_registry.DisabledPlugins {
		unloaded.WriteString("\t" + k + "\n")
	}
	bot.SendMessage(message.Chat, "Enabled plugins: ", nil)

	bot.SendMessage(message.Chat, loaded.String(), nil)
	bot.SendMessage(message.Chat, "Disabled plugins: ", nil)

	bot.SendMessage(message.Chat, unloaded.String(), nil)
}
