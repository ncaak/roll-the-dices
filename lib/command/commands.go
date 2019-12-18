package command

import (
//	"github.com/ncaak/roll-the-dices/lib/dice"
)

func NewTira(arg string) (c baseCommand) {
	c.resolve = aspectBasicRoll(arg, "1d20")
	return
}
