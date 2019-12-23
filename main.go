package main

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/command"
	"github.com/ncaak/roll-the-dices/lib/config"
	"github.com/ncaak/roll-the-dices/lib/dice"
	"github.com/ncaak/roll-the-dices/lib/request"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
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
			cmd, err := command.GetValidatedCommandOrError(res)
			if err != nil {
				log.Println("[WRN] " + err.Error())

			} else {
				cmd.Send(api)
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
