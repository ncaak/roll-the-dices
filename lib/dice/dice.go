package dice

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const MAX_REPETITIONS = 20

func initRoller(cmd string) Roller {
	return Roller{cmd, []check{}, []int{}, 0, ""}
}

// Main algorithm that goes through all the steps to retrieve args info, dice info, bonus info and tagging info
// Accepts a default Roll as argument to allow indirect rolls with only bonus and no dice
func resolve(command string, defaultRoll string) Roller {
	var r = initRoller(command)

	r.extractDice(defaultRoll)

	r.extractBonus()

	r.calcTotal()

	return r
}

// Regular roll returns a plain text reply
func Roll(command string, defaultRoll string) (string, error) {
	var r = resolve(command, defaultRoll)
	return r.GetReply(), nil
}

// Distribute roll returns a rich reply (currently markdown)
// It details separately rolls/bonus and supports multiple tags
// Used by commands: Agrupa
func Distribute(command string) (string, error) {
	var regexDices = regexp.MustCompile(`(?P<cmd>[^:]+)(?::(?P<tag>[^\s\+-]+))?`)
	var matches = regexDices.FindAllStringSubmatch(command, -1)
	var reply string
	if len(matches) == 0 {
		return errNotValidInput("dice.Distribute")
	}

	if len(matches) == 1 {
		reply = getNoTagsDist(command)

	} else {
		var str strings.Builder
		for _, match := range matches {
			var roll = getMapRegexGroups(regexDices.SubexpNames(), match)
			str.WriteString(getTaggedDistLine(roll))
		}
		reply = str.String()
	}

	return reply, nil
}

func getNoTagsDist(command string) string {
	var r = resolve(command, "1d20")
	r.setGroupedReply()
	return r.GetReply()
}

func getTaggedDistLine(roll map[string]string) string {
	var tag = "n/a"
	if roll["tag"] != "" {
		tag = roll["tag"]
	}

	roller := resolve(roll["cmd"], "")
	return roller.getDistReplyComp(tag)
}

// Returns a grouped similar rolls on a rich reply
// Used by commands: Repite
func Repeat(command string) (string, error) {
	var regexDices = regexp.MustCompile(`(?P<rpt>^\d+) (?P<cmd>.*)?`)
	var matches = regexDices.FindStringSubmatch(command)
	if len(matches) == 0 {
		return errNotValidInput("dice.Repeat")
	}

	var roll = getMapRegexGroups(regexDices.SubexpNames(), matches)
	var rpt = getRepetitions(roll["rpt"])
	var str strings.Builder

	for i := 0; i < rpt; i++ {
		roller := resolve(roll["cmd"], "1d20")
		str.WriteString(roller.getRepeatReplyComp())
	}
	return str.String(), nil
}

func getRepetitions(rpt string) int {
	var reps, _ = strconv.Atoi(rpt)
	// Avoid flooding with too many rolls
	if reps > MAX_REPETITIONS {
		return MAX_REPETITIONS
	}
	return reps
}

/*
 * Error typification
 */

func errNotValidInput(f string) (string, error) {
	return "", fmt.Errorf("Input was not valid for function: %s", f)
}
