package command

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
	"regexp"
	"strings"
)

/*
 * Interfaces to be used
 * Request interface communicates with API to send replies to the server
 * Msg interface communicates with structures needed to extract information to use with the API
 */
type Request interface {
	BasicReply(int, int, string)
}
type Source interface {
	GetChatId() int
	GetReplyId() int
}

/*
 *
 */
type baseCommand struct {
	source  Source
	resolve func() string
	send    func(Request, Source, string)
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

func (c baseCommand) Send(api Request) {
	var roll = c.resolve()

	c.send(api, c.source, roll)
}
