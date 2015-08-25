// Plugin register themselves here, the registry keep tracks of plugins to redirect the messages
package plugin_registry

import (
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"
	"log"
	"reflect"
	"strings"
)

type TelegramPlugin interface {
	Run(message telebot.Message)
	OnStart()
}

// These are are registered plugins
var Plugins = map[string]TelegramPlugin{}
var Commands = make(map[string]string)
var Config util.Config
var Bot *telebot.Bot

// Register a Plugin
func RegisterPlugin(p TelegramPlugin) {
	Plugins[KeyOf(p)] = p
}

// Remove a plugin
func RemovePlugin(plugin string) {
	delete(Plugins, plugin)
	log.Println(plugin + " removed from running plugins")
}

func KeyOf(p TelegramPlugin) string {
	return strings.TrimPrefix(reflect.TypeOf(p).String(), "*")
}

// Register a Command exported by a plugin
func RegisterCommand(command string, description string) {
	Commands[command] = description
}
