package command

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/dice"
	"regexp"
	"strings"
)

type baseCommand struct {
	result  string
	resolve func() dice.Roller
}

func validCommands() string {
	var VALID_COMMANDS = [...]string{"agrupa", "tira", "t", "v", "dv", "ayuda"}
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

func getValidatedCommandOrError(input string) (baseCommand, error) {
	var match = regexp.MustCompile(validCommands()).FindStringSubmatch(input)
	if len(match) == 0 {
		return baseCommand{}, fmt.Errorf("Received command is not valid")
	}

	return getCommand(match[1], match[2]), nil
}

func ResolveOrError(input string) (baseCommand, error) {
	// Command validation and structure initialization
	cmd, err := getValidatedCommandOrError(input)
	if err != nil {
		return cmd, err
	}

	roll := cmd.resolve()
	cmd.result = roll.GetReply()

	return cmd, err
}
