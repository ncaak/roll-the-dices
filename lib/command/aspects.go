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

func resolveDistRoll(input string) func() string {
	return func() string {
		var roller = dice.Distribute(input)
		return roller.GetReply()
	}
}

func resolveNoRoll() func() string {
	return func() string {
		return ""
	}
}

func sendBasicReply() func(Request, Source, string) {
	return func(api Request, source Source, roll string) {
		api.BasicReply(source.GetChatId(), source.GetReplyId(), roll)
	}
}

func sendKeyboard() func(Request, Source, string) {
	return func(api Request, source Source, _ string) {
		api.KeyboardReply(source.GetChatId(), source.GetReplyId())
	}
}

func sendMarkdownReply() func(Request, Source, string) {
	return func(api Request, source Source, roll string) {
		api.MarkdownReply(source.GetChatId(), source.GetReplyId(), roll)
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

func NewV(arg string) (c baseCommand) {
	c.resolve = resolveBasicRoll(arg, "h2d20")
	c.send = sendBasicReply()
	return
}

func NewDv(arg string) (c baseCommand) {
	c.resolve = resolveBasicRoll(arg, "l2d20")
	c.send = sendBasicReply()
	return
}

func NewT(arg string) (c baseCommand) {
	c.resolve = resolveNoRoll()
	c.send = sendKeyboard()
	return
}

func NewAgrupa(arg string) (c baseCommand) {
	c.resolve = resolveDistRoll(arg)
	c.send = sendMarkdownReply()
	return
}
