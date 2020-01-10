package command

import (
	"github.com/ncaak/roll-the-dices/lib/dice"
)

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
	c.resolve = func() (string, error) { return dice.Roll(arg, "1d20") }
	c.send = sendBasicReply()
	return
}

func NewV(arg string) (c baseCommand) {
	c.resolve = func() (string, error) { return dice.Roll(arg, "h2d20") }
	c.send = sendBasicReply()
	return
}

func NewDv(arg string) (c baseCommand) {
	c.resolve = func() (string, error) { return dice.Roll(arg, "l2d20") }
	c.send = sendBasicReply()
	return
}

func NewT(_ string) (c baseCommand) {
	c.resolve = func() (string, error) { return "", nil }
	c.send = sendKeyboard()
	return
}

func NewAgrupa(arg string) (c baseCommand) {
	c.resolve = func() (string, error) { return dice.Distribute(arg) }
	c.send = sendMarkdownReply()
	return
}

func NewAyuda(arg string) (c baseCommand) {
	c.resolve = func() (string, error) { return getHelp(arg), nil }
	c.send = sendMarkdownReply()
	return
}

func NewRepite(arg string) (c baseCommand) {
	c.resolve = func() (string, error) { return dice.Repeat(arg) }
	c.send = sendBasicReply()
	return
}