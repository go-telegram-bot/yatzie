package plugin_registry

import (
	"github.com/tucnak/telebot"
)

//Configuration needed for plugins and bot
type Config struct {
	Token         string
	CommandPrefix string
	Eloquens      bool
}

type TelegramPlugin interface {
	Run(bot *telebot.Bot, config Config, message telebot.Message)
}

// These are are registered plugins
var Plugins = []TelegramPlugin{}
var Commands = make(map[string]string)

// Register a Plugin
func RegisterPlugin(p TelegramPlugin) {
	Plugins = append(Plugins, p)
}

// Register a Command exported by a plugin
func RegisterCommand(command string, description string) {
	Commands[command] = description
}
