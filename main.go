package main

import (
	"fmt"
	C "github.com/ncaak/roll-the-dices/lib/command"
	"github.com/ncaak/roll-the-dices/lib/config"
	"github.com/ncaak/roll-the-dices/lib/dice"
	"github.com/ncaak/roll-the-dices/lib/request"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Println("[INF] Beginning routine")

	var settings = config.GetSettings()
	var db = storage.Init(settings.DataBase)
	var api = request.Init(settings.Api)
	var results = api.GetUpdates(db.GetOffset())

	defer db.Close()

	for _, res := range results {

		if res.IsCommand() == true {
//////
			cmd, err := C.GetValidatedCommandOrError(res)
			if err != nil {
				log.Println("[WRN] " + err.Error())

			} else {
				cmd.Send(api)
			}
//////
			var command = regexp.MustCompile(`/(agrupa|t|ayuda)(.*)`).FindStringSubmatch(res.GetCommand())

			if len(command) > 0 {
				var argument = strings.TrimSpace(command[2])
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
