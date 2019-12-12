package command

import (
	"fmt"
	"regexp"
	"strings"
)

func validCommands() string {
	var VALID_COMMANDS = [...]string{"agrupa", "tira", "t", "v", "dv", "ayuda"}

	return fmt.Sprintf("/(%s)(.*)", strings.Join(VALID_COMMANDS[:], "|"))
}

func getValidatedCommandOrError(seed string) (bool, error) {
	var command = regexp.MustCompile(validCommands()).FindStringSubmatch(seed)

	if len(command) > 0 {
		return true, nil
	}
	return false, fmt.Errorf("Received command is not valid")
}
