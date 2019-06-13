package main

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"github.com/ncaak/roll-the-dices/lib/conn"
	"github.com/ncaak/roll-the-dices/lib/dice"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"fmt"
	"regexp"
	"strings"
)

var acceptedCommands = [...]string{"tira","v","dv"}

func main() {
	log.Println("beginning routine")

	var env = "ENV_DEV"
	settings := config.GetSettings("ENV_DEV")
	var db = storage.Init(settings)
	var offset = db.GetOffset()
	var messages = conn.GetUpdates(env, offset)
	
	for _, msg := range messages {

		if msg.IsCommand() == true {
			var regexCommands = fmt.Sprintf("/(%s)(.*)", strings.Join(acceptedCommands[:], "|"))
			var command = regexp.MustCompile(regexCommands).FindStringSubmatch(msg.GetCommand())
			var reply string

			if len(command) > 0 {
				var argument = strings.TrimSpace(command[2])
				switch command[1] {
				case acceptedCommands[0]: // tira
					reply, _ = dice.Resolve(argument, "1d20")

				case acceptedCommands[1]: // v
					reply, _ = dice.Resolve(argument, "h2d20")

				case acceptedCommands[2]: // dv
					reply, _ = dice.Resolve(argument, "l2d20")
				}

				conn.SendReply(env, msg.Message.Chat.Id, reply, msg.Message.MessageId)
				fmt.Println("reply: ", reply)
			}

		}
	}

	if len(messages) > 0 {
		var newOffset = fmt.Sprintf("%d", messages[len(messages)-1].UpdateId +1)
		db.SetOffset(newOffset)
	}

	db.Close()
	log.Println("ending routine")
}
