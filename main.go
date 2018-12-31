package main

import (
	"github.com/ncaak/roll-the-dices/lib/conn"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"fmt"
)

func main() {
	log.Println("beginning routine")

	var offset = storage.GetUpdateOffset()
	var messages = conn.GetUpdates(offset)

	for _, msg := range messages {
		fmt.Println(msg)
		fmt.Println(msg.IsCommand())
	}

	storage.Close()

	log.Println("ending routine")
}
