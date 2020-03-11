package command

import (
	"fmt"
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
	CharKeyboardReply(int, int)
	DiceKeyboardReply(int, int)
	EditKeyboardReply(int, int, string)
	MarkdownReply(int, int, string)
}
type Source interface {
	GetChatId() int
	GetCommand() string
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
	var VALID_COMMANDS = [...]string{"tira", "v", "dv", "t", "agrupa", "ayuda", "repite", "pj"}
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
		cmd = NewT()
	case "agrupa":
		cmd = NewAgrupa(argument)
	case "ayuda":
		cmd = NewAyuda(argument)
	case "repite":
		cmd = NewRepite(argument)
	case "pj":
		cmd = NewPj()
	}
	return cmd
}

// Library main method to validate commands and initialize structure with required functionality
func GetValidatedCommandOrError(input Source) (baseCommand, error) {
	var match = regexp.MustCompile(validCommands()).FindStringSubmatch(input.GetCommand())
	if len(match) == 0 {
		return baseCommand{}, fmt.Errorf("Received command is not valid")
	}

	var command = getBaseCommand(match[1], match[2])
	command.source = input

	return command, nil
}

func SendErrorReply(input Source, api Request, err string) {
	var errCommand = NewError(err)
	errCommand.source = input
	errCommand.Run(api)
}

func SendCallbackReply(cb Source, api Request) {
	var cmd, _ = GetValidatedCommandOrError(cb)
	var roll, _ = cmd.resolve()

	api.EditKeyboardReply(cb.GetChatId(), cb.GetReplyId(), roll)
}
