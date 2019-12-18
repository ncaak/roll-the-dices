package command

import (
	"github.com/ncaak/roll-the-dices/lib/dice"
)

func aspectBasicRoll(input string, defaultRoll string) func() dice.Roller {
	return func() dice.Roller {
		return dice.Resolve(input, defaultRoll)
	}
}
