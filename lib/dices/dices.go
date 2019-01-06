package dices

import (
	"fmt"
	"regexp"
	"strconv"
	"math/rand"
	"time"
	"strings"
)


func Roll(roll string) string {
	fmt.Println("tirada: ", roll)
	
	var regexDices = regexp.MustCompile(`([+|-])?(\d)d(\d+)`)
	var total int
	var reply string
	
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

	

	if roll != "" {
		var regexBonus = regexp.MustCompile(`([+|-])?(\d+)`)

		for _, match := range regexBonus.FindAllStringSubmatch(roll, -1) {
			fmt.Println("bonus: ", match[0])
			bonus, _ := strconv.Atoi(match[0])
			total += bonus

			reply = fmt.Sprintf("%s %s", reply, match[0])
			roll = strings.Replace(roll, match[0], "", 1)
		}

	}

	if roll != "" {
		reply = fmt.Sprintf("%s:%s", strings.TrimSpace(roll), reply)
	}
			
	reply = fmt.Sprintf("%s = %d", reply, total)

	return reply
}
