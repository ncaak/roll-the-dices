package dices

import (
	"fmt"
	"regexp"
	"strconv"
	"math/rand"
	"time"
	"strings"
)

var reply string
var roll string
var total int

func reset(command string) {
	reply = ""
	roll = command
	total = 0
}

func rollDices(dices int, faces int) []int {
	var rollValues []int
	rand.Seed(time.Now().UnixNano())
	for rolls := 0; rolls < dices; rolls++ {
		rollValues = append(rollValues, rand.Intn(faces)+1)
	}

	return rollValues
}

func calcDices() {
	var regexDices = regexp.MustCompile(`([+|-])?(\d)d(\d+)`)

	for _, match := range regexDices.FindAllStringSubmatch(roll, -1) {
		var rollValues []int
		var dices, _ = strconv.Atoi(match[len(match)-2])
		var faces, _ = strconv.Atoi(match[len(match)-1])
		
		rollValues = rollDices(dices, faces)
		for _, value := range rollValues {
			total += value
		}

		reply = fmt.Sprintf("%s %s%d", reply, match[0], rollValues)
		roll = strings.Replace(roll, match[0], "", 1)
	}
}

func calcBonus() {
	var regexBonus = regexp.MustCompile(`([+|-])?(\d+)`)
	
	for _, match := range regexBonus.FindAllStringSubmatch(roll, -1) {
		bonus, _ := strconv.Atoi(match[0])
		total += bonus

		reply = fmt.Sprintf("%s %s", reply, match[0])
		roll = strings.Replace(roll, match[0], "", 1)
	}
}

func tag() {
	reply = fmt.Sprintf("%s: %s", strings.TrimSpace(roll), strings.TrimSpace(reply))
}

func Roll(command string) string {
	//fmt.Println("roll: ", command)
	if command != "" {
		reset(command)
	} else {
		reset("1d20")
	}

	if roll != "" {
		calcDices()
	}

	if roll != "" {
		calcBonus()
	}

	if roll != "" {
		tag()
	}
			
	reply = fmt.Sprintf("%s = %d", reply, total)
	
	return reply
}

func Advantage(command string) string {
	//fmt.Println("advantage: ", command)
	var rollValues []int
	
	reset(command)
	
	rollValues = rollDices(2,20)	
	if rollValues[0] > rollValues[1] {
		total = rollValues[0]
	} else {
		total = rollValues[1]
	}

	reply = fmt.Sprintf("2d20%d", rollValues)

	if roll != "" {
		calcDices()
	}
	
	if roll != "" {
		calcBonus()
	}

	if roll != "" {
		tag()
	}

	reply = fmt.Sprintf("%s = %d", reply, total)

	return reply
}

func Disadvantage(command string) string {
	//fmt.Println("disadvantage: ", command)
	var rollValues []int
	
	reset(command)
	
	rand.Seed(time.Now().UnixNano())
	for rolls := 0; rolls < 2; rolls++ {
		rollValues = append(rollValues, rand.Intn(20)+1)
	}

	if rollValues[0] < rollValues[1] {
		total = rollValues[0]
	} else {
		total = rollValues[1]
	}
	reply = fmt.Sprintf("2d20%d", rollValues)

	if roll != "" {
		calcBonus()
	}

	if roll != "" {
		tag()
	}

	reply = fmt.Sprintf("%s = %d", reply, total)

	return reply
}
