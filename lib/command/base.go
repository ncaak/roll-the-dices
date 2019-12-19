package command

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/request"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
	"regexp"
	"strings"
)

type baseCommand struct {
	source	structs.Message
	resolve   func() string
	send	func(request.Request, structs.Msg, string)
}

func validCommands() string {
	var VALID_COMMANDS = [...]string{"tira"}
	return fmt.Sprintf("/(%s)(.*)", strings.Join(VALID_COMMANDS[:], "|"))
}

func getCommand(inputCmd string, arg string) (cmd baseCommand) {
	var argument = strings.TrimSpace(arg)
	switch inputCmd {
	case "tira":
		cmd = NewTira(argument)

	}
	return cmd
}

func GetValidatedCommandOrError(input structs.Result) (baseCommand, error) {
	var match = regexp.MustCompile(validCommands()).FindStringSubmatch(input.GetCommand())
	if len(match) == 0 {
		return baseCommand{}, fmt.Errorf("Received command is not valid")
	}

	var command = getCommand(match[1], match[2])
	command.source = input.Message

	return command, nil
}


func (c baseCommand) Send (api request.Request) {
	var roll = c.resolve()

	c.send(api, c.source, roll)
}
