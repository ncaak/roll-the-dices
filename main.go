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

var acceptedCommands = [...]string{"tira", "v", "dv"}

func main() {
	log.Println("beginning routine")

	var settings = config.GetSettings(ENVIRONMENT)
	var db = storage.Init(settings.DataBase)
	var http = request.Init(settings.Api)
	var messages = http.GetUpdates(db.GetOffset())

	for _, msg := range messages {

		if msg.IsCommand() == true {
			var regexCommands = fmt.Sprintf("/(%s)(.*)", strings.Join(acceptedCommands[:], "|"))
			var command = regexp.MustCompile(regexCommands).FindStringSubmatch(msg.GetCommand())
			var roll dice.Roller

			if len(command) > 0 {
				var argument = strings.TrimSpace(command[2])
				switch command[1] {
				case acceptedCommands[0]: // tira
					roll = dice.Resolve(argument, "1d20")

				case acceptedCommands[1]: // v
					roll = dice.Resolve(argument, "h2d20")

				case acceptedCommands[2]: // dv
					roll = dice.Resolve(argument, "l2d20")
				}

				http.SendReply(msg, roll.FormatReply())
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
