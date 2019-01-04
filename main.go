package main

import (
	"github.com/ncaak/roll-the-dices/lib/conn"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"log"
	"fmt"
	"regexp"
	"strconv"
	"math/rand"
	"time"
	"strings"
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
				var roll = rollCmd[len(rollCmd)-1]
				fmt.Println("tirada: ", roll)
				var dicesRegex = regexp.MustCompile(`([+|-])?(\d)d(\d+)`)
			//	var dice = regexp.MustCompile(`([+|-])?(\d)d(\d+)`)



			//fmt.Println(dice.FindStringSubmatch(cmd))
			//fmt.Println(dice.FindString(cmd)
			var total = 0
			for i, match := range dicesRegex.FindAllStringSubmatch(roll, -1) {
				
				var rollValues []string
				fmt.Println(match, " encontrado ", i)

				var dices, _ = strconv.Atoi(match[len(match)-2])
				var faces, _ = strconv.Atoi(match[len(match)-1])
				roll = strings.Replace(roll, match[0], "", -1)

				rand.Seed(time.Now().Unix())
				for rolls := 0; rolls < dices; rolls++ {
					value := rand.Intn(faces) +1
					fmt.Println("ha salido un", value)
					total += value
//					strValue _ := strconv.Itoa(value)
					rollValues = append(rollValues, strconv.Itoa(value))
				}

				reply = fmt.Sprintf("%s %s%s", reply, match[0], rollValues)

				fmt.Println("resto de la tirada: ", roll)
			}
			//fmt.Println(regexp.MustCompile(`\/tira (.*)`).FindStringSubmatch(cmd))
				reply = fmt.Sprintf("%s= %d", reply, total)
	//		fmt.Println("Resultados ", rollValues)
				fmt.Println("Reply msg: ", reply)
				fmt.Println("Para un total de ", total)

			if roll != "" {
				fmt.Println("/// quedan numeros que sumar")
			}
			
			}
			//conn.SendReply(msg.Message.From.Id, reply, msg.Message.MessageId)
		}
	}

	storage.Close()

	log.Println("ending routine")
}
