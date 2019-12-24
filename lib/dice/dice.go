package dice

import (
	"fmt"
	"regexp"
	"strings"
)

func initRoller(cmd string) Roller {
	return Roller{cmd, []check{}, []int{}, 0, ""}
}

// Main algorithm that goes through all the steps to retrieve args info, dice info, bonus info and tagging info
// Accepts a default Roll as argument to allow indirect rolls with only bonus and no dice
func Resolve(command string, defaultRoll string) Roller {
	var r = initRoller(command)

	r.extractDice(defaultRoll)

	r.extractBonus()

	r.calcTotal()

	return r
}

// Distribute roll returns a rich reply (currently markdown)
// It details separately rolls/bonus and supports multiple tags
func Distribute(command string) (roller Roller) {
	var regexDices = regexp.MustCompile(`(?P<cmd>[^:]+)(?::(?P<key>[^\s\+-]+))?`)
	var matches = regexDices.FindAllStringSubmatch(command, -1)

	if len(matches) == 1 {
		roller = Resolve(command, "1d20")
		roller.setGroupedReply()

	} else {
		var str strings.Builder
		for _, match := range matches {
			var roll = getMapRegexGroups(regexDices.SubexpNames(), match)
			var tag = roll["key"]
			if tag == "" {
				tag = "n/a"
			}
			roller = Resolve(roll["cmd"], "")

			fmt.Fprintf(&str, "%s", roller.getDistributeReplyComponent(tag))
		}

		roller.reply = str.String()
	}

	return
}

func Repeat(command string) (Roller, error) {
	var regexDices = regexp.MustCompile(`(?P<rpt>^\d+) (?P<cmd>.*)?`)
	var matches = regexDices.FindStringSubmatch(command)

	if len(matches) == 0 {
		return Roller{}, fmt.Errorf("Input was not valid")
	}

	fmt.Println(getMapRegexGroups(regexDices.SubexpNames(), matches))

	return Roller{}, nil
}
