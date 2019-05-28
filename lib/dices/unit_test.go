package dices

import (
	"fmt"
	"testing"
)

// Combination of values die-face
type diceMatrix [][2]int

// Auxiliary function to sum up slice values
func sliceSum(slice []int) (total int) {
	for _, item := range slice {
		total += item
	}
	return
}

// Checks that rolls made in Roller contains:
// * the correct dice number
// * the correct faces number for the die
// * the dice results are inbounds
// * the total sum of each result
func checkRoll(t *testing.T, r Roller, matrix diceMatrix) {
	var checkTotal int

	for i, item := range r.checks {
		checkSum := sliceSum(item.results)
		dice, faces := matrix[i][0], matrix[i][1]
		// Checks the die number
		if item.dice != dice {
			t.Error(fmt.Sprintf("ERROR :: Wrong number of dices : Expected: %d, Got: %d", dice, item.dice))
		// Checks the die faces
		} else if item.faces != faces {
			t.Error(fmt.Sprintf("ERROR :: Wrong number of die faces : Expected: %d, Got: %d", faces, item.faces))
		// Checks the value of the check result is between possible minimum and maximum
		} else if checkSum < dice || checkSum > (dice * faces) {
			t.Error(fmt.Sprintf("ERROR :: Result outbounds dice possible values : Expected: %d-%d, Got: %d", dice, faces, checkSum))

		}
		checkTotal += checkSum
	}
	// Checks that check total coincides with check results sum
	if checkTotal != r.total {
		t.Error(fmt.Sprintf("ERROR :: Wrong check total : Expected: %d, Got: %d", checkTotal, r.total))
	}
}

// --- Tests ---

// Test 1d20 roll
// Minimum value: 1, Maximum value: 20
func TestBasicRoll(t *testing.T) {
	var test = "1d20"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '1d20[d1]= d1' d1 = [1-20]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}})
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test 2d20+1d6 roll
// Minimum value: 3, Maximum value: 46
func TestMultipleRoll(t *testing.T) {
	var test = "2d20+1d6"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d20[d1 d2]+1d6[d3]= d4' (d4=d1+d2+d3)")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,20},{1,6}})
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test 2d20+1d10+1d6 roll
// Minimum value: 4, Maximum value: 56
func TestMultipleRollWhitespaces(t *testing.T) {
	var test = "2d20 + 1d10 +1d6"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d20[d1 d2]+1d10[d3]+1d6[d4]= d5' (d5=d1+d2+d3+d4)")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,20},{1,10},{1,6}})
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test 1d20+5 roll
// Minimum value: 6, Maximum value: 25
func TestBonusBasicRoll(t *testing.T) {
	var test = "1d20+5"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '1d20[d1]+5= d2' d2 = [1-20]+5")
	var result, roller = Resolve(test)
	t.Log(roller)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}})
	t.Log(fmt.Sprintf("Result: %s", result))
}

