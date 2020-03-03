package command

import (
	"github.com/ncaak/roll-the-dices/lib/dice"
)

/*
 * Auxiliar functions to improve readability
 */
func returnStringOnResolve(message string) func() (string,error) {
	return func() (string, error) { return message, nil }
}
func returnRollOnResolve(s string, e error) func() (string,error) {
	return func() (string, error) { return s, e }
}

/*
 * Mutable functionalities
 * It adds polymorphism for the different commands
 */

// Replies to API functions
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
	c.resolve = returnRollOnResolve(dice.Roll(arg, "1d20"))
	c.send = sendBasicReply()
	return
}

func NewV(arg string) (c baseCommand) {
	c.resolve = returnRollOnResolve(dice.Roll(arg, "h2d20"))
	c.send = sendBasicReply()
	return
}

func NewDv(arg string) (c baseCommand) {
	c.resolve = returnRollOnResolve(dice.Roll(arg, "l2d20"))
	c.send = sendBasicReply()
	return
}

func NewT(_ string) (c baseCommand) {
	c.resolve = returnStringOnResolve("")
	c.send = sendKeyboard()
	return
}

func NewAgrupa(arg string) (c baseCommand) {
	c.resolve = returnRollOnResolve(dice.Distribute(arg))
	c.send = sendMarkdownReply()
	return
}

func NewAyuda(arg string) (c baseCommand) {
	c.resolve = returnStringOnResolve(getHelp(arg))
	c.send = sendMarkdownReply()
	return
}

func NewRepite(arg string) (c baseCommand) {
	c.resolve = returnRollOnResolve(dice.Repeat(arg))
	c.send = sendBasicReply()
	return
}

func NewError(message string) (c baseCommand) {
	c.resolve = returnStringOnResolve(message)
	c.send = sendMarkdownReply()
	return
}
