package main

import (
	"github.com/ncaak/roll-the-dices/lib/conn"
	"github.com/ncaak/roll-the-dices/lib/dices"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"fmt"
	"regexp"
	"strings"
)

var acceptedCommands = [...]string{"tira","v","dv"}

func main() {
	log.Println("beginning routine")

	var offset = storage.GetUpdateOffset()
	var messages = conn.GetUpdates(offset)

	for _, msg := range messages {

		if msg.IsCommand() == true {
			var regexCommands = fmt.Sprintf("/(%s)(.*)", strings.Join(acceptedCommands[:], "|"))
			var command = regexp.MustCompile(regexCommands).FindStringSubmatch(msg.GetCommand())
			var reply string

			if len(command) > 0 {
				var argument = strings.TrimSpace(command[2])
				switch command[1] {
				case acceptedCommands[0]:
					reply = dices.Roll(argument)

				case acceptedCommands[1]:
					reply = dices.Advantage(argument)

				case acceptedCommands[2]:
					fmt.Println("DesVentiaja", command[2])
				}

				//conn.SendReply(msg.Message.Chat.Id, reply, msg.Message.MessageId)
				fmt.Println("reply: ", reply)
			}

		}
	}

	if len(messages) > 0 {
		//var newOffset = fmt.Sprintf("%d", messages[len(messages)-1].UpdateId +1)
		//storage.SetLastUpdateId(newOffset)
		fmt.Println("debug mode active")
	}
	storage.Close()

	log.Println("ending routine")
}
