package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type check struct {
	dice    int
	faces   int
	results []int
	total   int
}

type Roller struct {
	command string
	checks  []check
	bonus   []int
	total   int
	reply   string
}

// Sums up every roll made with dices and modifier included as a bonus (negative or positive)
func (r *Roller) calcTotal() {
	// Searches for every result of every check
	for _, check := range r.checks {
		r.total += check.total
	}
	// Searches for every bonus
	for _, mod := range r.bonus {
		r.total += mod
	}
}

// Extracts strings referring dice rolls from the command and add them as check items in a slice
// Accepts a default Roll as argument to allow indirect rolls with only bonus and no dice
func (r *Roller) extractDice(defaultRoll string) {
	var regexDices = regexp.MustCompile(`[ ]*\+?[ ]*(?P<arg>\d)?(?P<mod>[hl])?(?P<dice>\d+)d(?P<faces>\d+)`)
	var matches = regexDices.FindAllStringSubmatch(r.command, -1)
	// In case no die was found insert the defaultRoll as a pre-generated roll
	if len(matches) == 0 && defaultRoll != "" {
		r.command = fmt.Sprintf("%s%s", defaultRoll, r.command)
	}

	for _, match := range matches {
		var roll = getMapRegexGroups(regexDices.SubexpNames(), match)
		var dices, _ = strconv.Atoi(roll["dice"])
		var faces, _ = strconv.Atoi(roll["faces"])
		r.newCheck(dices, faces, getCheckAction(roll["mod"], roll["arg"]))
		r.extractFromCommand(roll["command"])
	}
}

// Extracts strings referring bonuses from the command and add them as bonus in a slice
func (r *Roller) extractBonus() {
	var regexBonus = regexp.MustCompile(`(\s?[+|-]\s?)(\d+)`)

	for _, match := range regexBonus.FindAllStringSubmatch(r.command, -1) {
		// index 1 gets the symbol '+' or '-' with possible whitespaces
		// index 2 gets the absolute value of the bonus
		var strBonus = fmt.Sprintf("%s%s", strings.TrimSpace(match[1]), match[2])
		var bonus, _ = strconv.Atoi(strBonus)
		r.bonus = append(r.bonus, bonus)
		// index 0 gets the complete roll, including possible whitespaces between symbol and value
		r.extractFromCommand(match[0])
	}
}

// Removes the string executed from the command to avoid false positives in regex
// e.g. dice could be catched as bonuses leaving a 'dx'
func (r *Roller) extractFromCommand(usedOrder string) {
	r.command = strings.Replace(r.command, usedOrder, "", 1)
}

// Generates results with given dices and die faces
func (r *Roller) newCheck(dice int, faces int, action func([]int) int) {
	var results = []int{}

	for rolls := 0; rolls < dice; rolls++ {
		// Generating new seed every execution
		rand.Seed(time.Now().UnixNano())
		// Minimum die value is 1, in randomizer is 0
		var roll = rand.Intn(faces) + 1
		// Retrieves the action value, by default is the sum of each value but can be modified
		results = append(results, roll)
	}
	var total = action(results)
	// new structure for each check
	r.checks = append(r.checks, check{dice, faces, results, total})
}

// Provides polymorphism to newCheck method, returning an action depending on the modifier
// Valid modifiers that can affect a roll (h, l)
func getCheckAction(modifier string, argument string) (action func([]int) int) {
	var arg, _ = strconv.Atoi(argument)

	switch modifier {
	case "h": // Higher roll on a check
		action = func(results []int) int {
			// Avoid wrong command spelling
			if arg > len(results) {
				arg = len(results)
			} else if arg == 0 {
				arg = 1
			}
			// Sorts the results by increasing order
			tmp := make([]int, len(results))
			copy(tmp, results)
			sort.Ints(tmp)
			// Returns slice with the number of items defined by argument (or 1)
			return getSliceItemsTotal(tmp[len(tmp)-arg:])
		}

	case "l": // Lower roll on a check
		action = func(results []int) int {
			sort.Ints(results)
			return results[0]
		}

	default: // default action is rolls sum
		action = func(results []int) (total int) {
			for _, value := range results {
				total += value
			}
			return
		}
	}
	return
}
