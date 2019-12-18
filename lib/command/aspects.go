package command

import (
	"github.com/ncaak/roll-the-dices/lib/dice"
)

func aspectBasicRoll(input string, defaultRoll string) (func() string) {
	return func() string {
		var roller = dice.Resolve(input, defaultRoll)
		return roller.GetReply()
	}
}
