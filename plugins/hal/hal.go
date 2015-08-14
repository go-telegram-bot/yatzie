package hal

import (
	"github.com/freahs/microhal"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	"os"
	"strings"
	"time"
)

var quips = []string{
	"FOR SCIENCE!",
	"because... reasons.",
	"it's super effective!",
	"because... why not?",
	"was it good for you?",
	"given the alternative, yep, worth it!",
	"don't ask...",
	"then makes a sandwich.",
	"oh noes!",
	"did I do that?",
	"why must you turn this place into a house of lies!",
	"really???",
	"LLLLEEEEEERRRRRROOOOYYYY JEEEENNNKINNNS!",
	"DOH!",
	"Giggity!",
}

const MarkovChainOrder int = 3

type MyPlugin struct {
	brainIn  chan<- string
	brainOut <-chan string
}

func (m *MyPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
	if !strings.HasPrefix(message.Text, config.CommandPrefix) && !util.MatchAnyURL(message.Text) {
		text := strings.Replace(message.Text, "@"+bot.Identity.Username, "", -1)
		// then call hal for random answers
		if len(text) >= MarkovChainOrder {
			m.brainIn <- text
			res := <-m.brainOut
			if res == "" && config.Eloquens == true {
				bot.SendMessage(message.Chat, util.RandomFromArray(quips), nil)

			} else {
				bot.SendMessage(message.Chat,
					res, nil)

			}
		} else {
			bot.SendMessage(message.Chat, util.RandomFromArray(quips), nil)
		}

	}
}

func init() {
	brainFile := "Brain"
	var brain *microhal.Microhal

	if _, err := os.Stat(brainFile); os.IsNotExist(err) {
		brain = microhal.NewMicrohal(brainFile, MarkovChainOrder)
	} else {
		brain = microhal.LoadMicrohal(brainFile)
	}
	brainIn, brainOut := brain.Start(10000*time.Millisecond, 250)
	my := &MyPlugin{brainIn: brainIn, brainOut: brainOut}
	plugin_registry.RegisterPlugin(my)
}
