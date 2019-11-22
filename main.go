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
	log.Println("[INF] Beginning routine")

	var settings = config.GetSettings(ENVIRONMENT)
	var db = storage.Init(settings.DataBase)
	var api = request.Init(settings.Api)
	var results = api.GetUpdates(db.GetOffset())

	defer db.Close()

	for _, res := range results {

		if res.IsCommand() == true {
			var command = regexp.MustCompile(`/(agrupa|tira|t|v|dv|ayuda)(.*)`).FindStringSubmatch(res.GetCommand())

			if len(command) > 0 {
				var argument = strings.TrimSpace(command[2])
				var rollCommands = map[string]string{
					"tira": "1d20",
					"v":    "h2d20",
					"dv":   "l2d20",
				}

				// Detects the command entered being a roll command
				if defRoll := rollCommands[command[1]]; defRoll != "" {
					var roll = dice.Resolve(argument, defRoll)
					api.Reply(res.Message, roll.GetReply())
					fmt.Println("reply provided")

				} else {
					switch command[1] {
					case "agrupa":
						var roll = dice.Distribute(argument)
						api.ReplyMarkdown(res.Message, roll.GetReply())
						fmt.Println("rich reply provided")

					case "t":
						api.ReplyInlineKeyboard(res.Message)
						fmt.Println("inline keyboard provided")

					case "ayuda":
						api.ReplyMarkdown(res.Message, dice.HELP)
						fmt.Println("help provided")
					}
				}
			}
		} else if res.IsCallback() {
			// A callback is triggered when someone clicks an inline keyboard
			var roll = dice.Resolve(res.Callback.Data, "1d20")
			api.EditKeyboardReply(res.Callback, roll.GetReply())
		}
	}

	if len(results) > 0 {
		var newOffset = fmt.Sprintf("%d", results[len(results)-1].UpdateId+1)
		db.SetOffset(newOffset)
	}

	log.Println("[INF] Ending routine")
}
