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

func calcDices() {
	var regexDices = regexp.MustCompile(`([+|-])?(\d)d(\d+)`)

	for _, match := range regexDices.FindAllStringSubmatch(roll, -1) {
		var rollValues []string
		var dices, _ = strconv.Atoi(match[len(match)-2])
		var faces, _ = strconv.Atoi(match[len(match)-1])

		rand.Seed(time.Now().UnixNano())
		for rolls := 0; rolls < dices; rolls++ {
			result := rand.Intn(faces) +1
			if match[1] == "-" {
				result = result * -1
			}

			total += result
			rollValues = append(rollValues, strconv.Itoa(result))
		}
		
		reply = fmt.Sprintf("%s %s%s", reply, match[0], rollValues)
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

	calcDices()

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
	
	rand.Seed(time.Now().UnixNano())
	for rolls := 0; rolls < 2; rolls++ {
		rollValues = append(rollValues, rand.Intn(20)+1)
	}

	if rollValues[0] > rollValues[1] {
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
