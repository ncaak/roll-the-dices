package dices

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type check struct {
	dice int
	faces int
	results []int
}

type Roller struct {
	command string
	checks []check
	bonus []int
	total int
}

func Resolve(command string) (string, Roller) {
	var r = Roller{command, []check{}, []int{},  0}

	r.extractDice()

	r.extractBonus()

	r.calcTotal()

	return r.formatReply(), r
}

// Sums up every roll made with dices and modifier included as a bonus (negative or positive)
func (r *Roller) calcTotal () {
	// Searches for every result of every check
	for _, check := range r.checks {
		for _, result := range check.results {
			r.total += result
		}
	}
	// Searches for every bonus
	for _, mod := range r.bonus {
		r.total += mod
	}
}

func (r *Roller) extractDice () {
	var regexDices = regexp.MustCompile(`(\s?\+?\s?)?(\d)d(\d+)`)

	for _, match := range regexDices.FindAllStringSubmatch(r.command, -1) {
		var dices, _ = strconv.Atoi(match[len(match)-2])
		var faces, _ = strconv.Atoi(match[len(match)-1])
		r.newCheck(dices, faces)
		r.extractFromCommand(match[0])
	}
}

func (r *Roller) extractBonus () {
	var regexBonus = regexp.MustCompile(`(\s?[+|-]\s?)(\d+)`)

	for _, match := range regexBonus.FindAllStringSubmatch(r.command, -1) {
		var strBonus = fmt.Sprintf("%s%s", strings.TrimSpace(match[1]), match[2])
		var bonus, _ = strconv.Atoi(strBonus)
		r.bonus = append(r.bonus, bonus)
		r.extractFromCommand(match[0])
	}
}

func (r *Roller) extractFromCommand (usedOrder string) {
	r.command = strings.Replace(r.command, usedOrder, "", 1)
}

// Generates a String with a verbose result of the roll
// * Retrieves every result of every check done
// * Retrieves every bonus
func (r *Roller) formatReply () string {
	var fmtReply strings.Builder
	// Finds every check and results to write it verbosely
	for index, check := range r.checks {
		if index > 0 {
			fmtReply.WriteString("+")
		}
		fmtReply.WriteString(fmt.Sprintf("%dd%d%d", check.dice, check.faces, check.results))

	}
	// Finds every bonus and writes it after the dice
	for _, bonus := range r.bonus {
		if bonus > 0 {
			fmtReply.WriteString("+")
		}
		fmtReply.WriteString(fmt.Sprintf("%d", bonus))
	}
	// Append equals symbol and the total sum of the roll
	fmtReply.WriteString(fmt.Sprintf("= %d", r.total))

	return fmtReply.String()
}

func (r *Roller) newCheck (dice int, faces int) {
	var ch = check{dice, faces, []int{}}
	rand.Seed(time.Now().UnixNano())
	for rolls := 0; rolls < dice; rolls++ {
		ch.results = append(ch.results, rand.Intn(faces)+1)
	}

	r.checks = append(r.checks, ch)
}
