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

package admin

import (
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/inconshreveable/go-update"

	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"gopkg.in/tucnak/telebot.v1"
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

	if strings.Contains(message.Text, config.CommandPrefix+"update") {
		url := util.StripPluginCommand(message.Text, config.CommandPrefix, "update")
		if url != "" {
			bot.SendMessage(message.Chat, "Upgrading with "+url, nil)
			err := doUpdate(url)
			if err != nil {
				bot.SendMessage(message.Chat, err.Error(), nil)
			} else {
				bot.SendMessage(message.Chat, "Everything went OK :)", nil)
				ForkExec()
				bot.SendMessage(message.Chat, "If all went straight you should see me again", nil)
				os.Exit(0)
			}
		}
	}
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	return err
}

func lookPath() (argv0 string, err error) {
	argv0, err = exec.LookPath(os.Args[0])
	if nil != err {
		return
	}
	if _, err = os.Stat(argv0); nil != err {
		return
	}
	return
}

func ForkExec() error {
	argv0, err := lookPath()
	if nil != err {
		return err
	}
	wd, err := os.Getwd()
	if nil != err {
		return err
	}

	p, err := os.StartProcess(argv0, os.Args, &os.ProcAttr{
		Dir: wd,
		Sys: &syscall.SysProcAttr{},
	})
	if nil != err {
		return err
	}
	log.Println("spawned child", p.Pid)

	return nil
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
