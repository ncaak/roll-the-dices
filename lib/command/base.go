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
	KeyboardReply(int, int)
	MarkdownReply(int, int, string)
}
type Source interface {
	GetChatId() int
	GetReplyId() int
}

/*
 * Structure used to handle commands
 * resolve and send methods are mutable and they depend on structure initializers
 */
type baseCommand struct {
	source  Source
	resolve func() (string, error)
	send    func(Request, Source, string)
}

// Unique entry point to handle every command
func (c baseCommand) Run(api Request) error {
	var roll, err = c.resolve()
	if err == nil {
		c.send(api, c.source, roll)
	}
	return err
}

/*
 * Common library functions
 * Validators and Initializers orchestrator
 */
func validCommands() string {
	var VALID_COMMANDS = [...]string{"tira", "v", "dv", "t", "agrupa", "ayuda", "repite"}
	return fmt.Sprintf(`/(%s)(\s.*)?$`, strings.Join(VALID_COMMANDS[:], "|"))
}

func getBaseCommand(inputCmd string, arg string) (cmd baseCommand) {
	var argument = strings.TrimSpace(arg)
	switch inputCmd {
	case "tira":
		cmd = NewTira(argument)
	case "v":
		cmd = NewV(argument)
	case "dv":
		cmd = NewDv(argument)
	case "t":
		cmd = NewT(argument)
	case "agrupa":
		cmd = NewAgrupa(argument)
	case "ayuda":
		cmd = NewAyuda(argument)
	case "repite":
		cmd = NewRepite(argument)
	}
	return cmd
}

// Library main method to validate commands and initialize structure with required functionality
func GetValidatedCommandOrError(input structs.Result) (baseCommand, error) {
	var match = regexp.MustCompile(validCommands()).FindStringSubmatch(input.GetCommand())
	if len(match) == 0 {
		return baseCommand{}, fmt.Errorf("Received command is not valid")
	}
	var command = getBaseCommand(match[1], match[2])
	command.source = input.Message

	return command, nil
}
