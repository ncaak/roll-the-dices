package main

import (
	"github.com/ncaak/roll-the-dices/lib/conn"
	"github.com/ncaak/roll-the-dices/lib/dices"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"fmt"
	"regexp"
)

func main() {
	log.Println("beginning routine")

	var offset = storage.GetUpdateOffset()
	var messages = conn.GetUpdates(offset)

	for _, msg := range messages {
		fmt.Println(msg)
		fmt.Println(msg.IsCommand())

		if msg.IsCommand() == true {
			var cmd = msg.Message.Text
			var rollCmd = regexp.MustCompile(`\/tira (.*)`).FindStringSubmatch(cmd)
			var reply string

			if len(rollCmd) > 0 {
				var rollString = rollCmd[len(rollCmd)-1]
				reply = dices.Roll(rollString)
			}
			
			conn.SendReply(msg.Message.From.Id, reply, msg.Message.MessageId)
		}
	}

	storage.Close()

	log.Println("ending routine")
}
