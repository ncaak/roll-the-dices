package command

import (
	"github.com/ncaak/roll-the-dices/lib/dice"
)

/*
 * Mutable functionalities
 * It adds polymorphism for the different commands
 */

func resolveBasicRoll(input string, defaultRoll string) func() string {
	return func() string {
		var roller = dice.Resolve(input, defaultRoll)
		return roller.GetReply()
	}
}

func sendBasicReply() func(Request, Source, string) {
	return func(api Request, source Source, roll string) {
		api.BasicReply(source.GetChatId(), source.GetReplyId(), roll)
	}
}

/*
 * Commands initializers
 * Each one sets the functions needed for the command to work
 */

func NewTira(arg string) (c baseCommand) {
	c.resolve = resolveBasicRoll(arg, "1d20")
	c.send = sendBasicReply()
	return
}
