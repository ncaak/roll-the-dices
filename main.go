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

const COMMAND_UNKNOWN = "El comando recibido es erróneo.\nCompruebe la /ayuda"
const PARAMETERS_ERROR = "El comando no acepta los datos introducidos.\nCompruebe la /ayuda"

func panicOnError(err error) {
	if err != nil {
		log.Printf("[ERR] %s", err.Error())
		panic(err)
	}
}

func main() {
	log.Println("[INF] Beginning routine")

	var settings, errConf = config.GetSettings()
	panicOnError(errConf)

	var db, errDB = storage.Init(settings)
	panicOnError(errDB)
	defer db.Close()

	var api = request.Init(settings)
	var updates = api.GetUpdates(db.GetOffset())

	for _, update := range updates {

		if update.IsCommand() {
			cmd, err := command.GetValidatedCommandOrError(update.Message)
			if err != nil {
				log.Println("[WRN] " + err.Error())
				command.SendErrorReply(update.Message, api, COMMAND_UNKNOWN)
				continue
			}
			
			errOnRun := cmd.Run(api)
			if errOnRun != nil {
				log.Println("[ERR] " + errOnRun.Error())
				command.SendErrorReply(update.Message, api, PARAMETERS_ERROR)
			}

		} else if update.IsCallback() {
			// A callback is triggered when someone clicks an inline keyboard
			var roll, _ = dice.Roll(update.Callback.Data, "1d20")
			api.EditKeyboardReply(update.Callback, roll)
		}
	}

	if len(updates) > 0 {
		var newOffset = fmt.Sprintf("%d", updates[len(updates)-1].UpdateId+1)
		db.SetOffset(newOffset)
	}

	log.Println("[INF] Ending routine")
}
