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
	OnStop()
}

// These are are registered plugins
var Plugins = map[string]TelegramPlugin{}
var DisabledPlugins = map[string]TelegramPlugin{}
var Commands = make(map[string]string)
var Config util.Config
var Bot *telebot.Bot

// Register a Plugin
func RegisterPlugin(p TelegramPlugin) {
	Plugins[KeyOf(p)] = p
}

// Disable a plugin
func DisablePlugin(plugin string) bool {
	plugin = strings.TrimSpace(plugin)
	DisabledPlugins[plugin] = Plugins[plugin]
	_, disabled := DisabledPlugins[plugin]
	if disabled {
		delete(Plugins, plugin)
		DisabledPlugins[plugin].OnStop()

		log.Println(plugin + " removed from running plugins")
	}

	return disabled
}

// Enable a plugin
func EnablePlugin(plugin string) bool {
	plugin = strings.TrimSpace(plugin)

	_, PluginExists := Plugins[plugin]
	if PluginExists {
		return true
	}

	PluginInstance, InstanceExists := DisabledPlugins[plugin]
	Plugins[plugin] = PluginInstance
	if InstanceExists {

		delete(DisabledPlugins, plugin)
		PluginInstance.OnStart()

		log.Println(plugin + " enabled ")
		return true
	}
	return false
}

func KeyOf(p TelegramPlugin) string {
	return strings.TrimPrefix(reflect.TypeOf(p).String(), "*")
}

// Register a Command exported by a plugin
func RegisterCommand(command string, description string) {
	Commands[command] = description
}

// UnRegister a Command exported by a plugin
func UnregisterCommand(command string) {
	delete(Commands, command)
}
