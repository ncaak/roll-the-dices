package command

import (
	"github.com/ncaak/roll-the-dices/lib/dice"
	"github.com/ncaak/roll-the-dices/lib/request"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
)

func resolveBasicRoll(input string, defaultRoll string) func() string {
	return func() string {
		var roller = dice.Resolve(input, defaultRoll)
		return roller.GetReply()
	}
}

func sendBasicReply() (func (request.Request, structs.Msg, string)) {
	return func (api request.Request, source structs.Msg, roll string) {
		api.BasicReply(source, roll)
	}
}
