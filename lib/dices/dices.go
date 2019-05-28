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

	r.extractModifiers()

	r.calcTotal()

	return r.formatReply(), r
}

func (r *Roller) calcTotal () {
	for _, check := range r.checks {
		for _, result := range check.results {
			r.total += result
		}
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

func (r *Roller) extractModifiers () {
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

func (r *Roller) formatReply () string {
	var fmtReply strings.Builder
	for index, check := range r.checks {
		var str = fmt.Sprintf("%dd%d%d", check.dice, check.faces, check.results)
		if index > 0 {
			fmtReply.WriteString("+")
		}
		fmtReply.WriteString(str)
	}
	
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
