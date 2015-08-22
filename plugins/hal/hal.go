// The hal plugin uses microhal to give the bot a little bit of soul
// If mentioned or talked directly, this plugin make your bot answer
package hal

import (
	"github.com/freahs/microhal"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	"log"
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

type HalPlugin struct {
	brainIn          chan<- string
	brainOut         <-chan string
	started          bool
	MarkovChainOrder int
}

func (m *HalPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config
	if !strings.HasPrefix(message.Text, config.CommandPrefix) && !util.MatchAnyURL(message.Text) {
		if m.started == false {
			loadHAL(m)
		}
		text := strings.Replace(message.Text, "@"+bot.Identity.Username, "", -1)
		// then call hal for random answers
		if len(text) > m.MarkovChainOrder {
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

func loadHAL(m *HalPlugin) {
	brainFile := "Brain"
	m.MarkovChainOrder = 3 // our default
	if plugin_registry.Config.HALBrainfile != "" {
		brainFile = plugin_registry.Config.HALBrainfile
	}
	if plugin_registry.Config.HALMarkovChainOrder != 0 {
		m.MarkovChainOrder = plugin_registry.Config.HALMarkovChainOrder
	}
	var brain *microhal.Microhal
	log.Println("My brainfile is:" + brainFile)
	if _, err := os.Stat(brainFile); os.IsNotExist(err) {
		brain = microhal.NewMicrohal(brainFile, m.MarkovChainOrder)
	} else {
		brain = microhal.LoadMicrohal(brainFile)
	}
	brainIn, brainOut := brain.Start(10000*time.Millisecond, 250)
	m.brainIn = brainIn
	m.brainOut = brainOut
	m.started = true
}

func init() {
	plugin_registry.RegisterPlugin(&HalPlugin{})
}
