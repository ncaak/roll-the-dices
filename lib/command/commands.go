package command

import (
//	"github.com/ncaak/roll-the-dices/lib/dice"
)

func NewTira(arg string) (c baseCommand) {
	c.resolve = resolveBasicRoll(arg, "1d20")
	c.send = sendBasicReply()
	return
}
