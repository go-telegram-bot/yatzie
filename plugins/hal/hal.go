// The hal plugin uses microhal to give the bot a little bit of soul
// If mentioned or talked directly, this plugin make your bot answer
package hal

import (
	cobe "github.com/mudler/go.cobe"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"
	"github.com/tucnak/telebot"

	"log"
	"strings"
)


type HalPlugin struct {
	Brain *cobe.Cobe2Brain
}

func init() {
	plugin_registry.RegisterPlugin(&HalPlugin{})
}

func (m *HalPlugin) OnStart() {
	b, err := cobe.OpenCobe2Brain(plugin_registry.Config.BrainFile)
	m.Brain = b
	if err != nil {
		log.Fatalf("Opening brain file: %s", err)
	}
	log.Println("[HalPlugin] Started")
}

func (m *HalPlugin) OnStop() {
		log.Println("[HalPlugin] Disabled")
}

func (m *HalPlugin) Run(message telebot.Message) {
	bot := plugin_registry.Bot
	config := plugin_registry.Config

	if !strings.HasPrefix(message.Text, config.CommandPrefix) && !util.MatchAnyURL(message.Text) {
		text := strings.Replace(message.Text, "@"+bot.Identity.Username, "", -1)
		m.Brain.Learn(text)
		if message.IsPersonal() || strings.Contains(message.Text,"@"+bot.Identity.Username) {
			bot.SendMessage(message.Chat, m.Brain.Reply(text), nil)
		}

	}
}
