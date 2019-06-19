package main

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/config"
	"github.com/ncaak/roll-the-dices/lib/dice"
	"github.com/ncaak/roll-the-dices/lib/request"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"regexp"
	"strings"
)

const ENVIRONMENT = "ENV_DEV"

func main() {
	log.Println("beginning routine")

	var settings = config.GetSettings(ENVIRONMENT)
	var db = storage.Init(settings.DataBase)
	var api = request.Init(settings.Api)
	var messages = api.GetUpdates(db.GetOffset())

	for _, msg := range messages {

		if msg.IsCommand() == true {
			var command = regexp.MustCompile(`/(tira|v|dv)(.*)`).FindStringSubmatch(msg.GetCommand())

			if len(command) > 0 {
				var argument = strings.TrimSpace(command[2])
				var defaultValues = map[string]string{
					"tira": "1d20",
					"v":    "h2d20",
					"dv":   "l2d20",
				}
				var roll = dice.Resolve(argument, defaultValues[command[1]])

				api.SendReply(msg, roll.FormatReply())
				fmt.Println("reply: ", roll.FormatReply())
			}

		}
	}

	if len(messages) > 0 {
		var newOffset = fmt.Sprintf("%d", messages[len(messages)-1].UpdateId+1)
		db.SetOffset(newOffset)
	}

	db.Close()
	log.Println("ending routine")
}
