package dice

import (
	"fmt"
	"testing"
	"strings"
)

// Combination of values die-face
type diceMatrix [][2]int

// --- Auxiliary functions ---

// Sum up slice values iterating through the slice
func sliceSum() func([]int) int {
	return func(slice []int) (total int) {
		for _, item := range slice {
			total += item
		}
		return
	}
}

// Returns the higher value among the slice items
func sliceTop() func([]int) int {
	return func(slice []int) (top int) {
		for _, item := range slice {
			if item > top {
				top = item
			}
		}
		return
	}
}

// Checks that rolls made in Roller contains:
// * the correct dice number
// * the correct faces number for the die
// * the dice results are inbounds
// * the total sum of each result
func checkRoll(t *testing.T, r Roller, matrix diceMatrix, bonus []int, action func([]int) int) {
	var checkTotal int

	for i, item := range r.checks {
		var checkSum = action(item.results)
		var dice, faces = matrix[i][0], matrix[i][1]
		// Checks the die number
		if item.dice != dice {
			t.Error(fmt.Sprintf("ERROR :: Wrong number of dices : Expected: %d, Got: %d", dice, item.dice))
		// Checks the die faces
		} else if item.faces != faces {
			t.Error(fmt.Sprintf("ERROR :: Wrong number of die faces : Expected: %d, Got: %d", faces, item.faces))
		// Checks the value of the check result is between possible minimum and maximum
		} else if checkSum < dice || checkSum > (dice * faces) {
			t.Error(fmt.Sprintf(
				"ERROR :: Result outbounds dice possible values : Expected: %d-%d, Got: %d",
				dice,
				dice * faces,
				checkSum,
			))
		}
		checkTotal += checkSum
	}
	// Checks that check total coincides with check results sum
	if checkTotal + action(bonus) != r.total {
		t.Error(fmt.Sprintf("ERROR :: Wrong check total : Expected: %d, Got: %d", checkTotal, r.total))
	}
}

// --- Tests ---

// Test 1d20 roll
// Minimum value: 1, Maximum value: 20
func TestRollBasic(t *testing.T) {
	var test = "1d20"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '1d20[d1]= d1' d1 = [1-20]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}}, []int{}, sliceSum())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test 2d20+1d6 roll
// Minimum value: 3, Maximum value: 46
func TestRollMultiple(t *testing.T) {
	var test = "2d20+1d6"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d20[d1 d2]+1d6[d3]= d4' (d4=d1+d2+d3)")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,20},{1,6}}, []int{}, sliceSum())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test 2d20+1d10+1d6 roll
// Minimum value: 4, Maximum value: 56
func TestRollWhitespaces(t *testing.T) {
	var test = "2d20 + 1d10 +1d6"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d20[d1 d2]+1d10[d3]+1d6[d4]= d5' (d5=d1+d2+d3+d4)")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,20},{1,10},{1,6}}, []int{}, sliceSum())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test flat +5 bonus
// Minimum value: 6, Maximum value: 25
func TestBonusBasic(t *testing.T) {
	var test = "1d20+5"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '1d20[d1]+5= d2' d2 = [6-25]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}}, []int{5}, sliceSum())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test flat -3 bonus
// Minimum value -1, Maximum value 17
func TestBonusNegative(t *testing.T) {
	var test = "2d10-3"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d10[d1]-3= d2' d2 = [-1-17]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,10}}, []int{-3}, sliceSum())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test multiple bonus
// Minimum value 5, Maximum value 24
func TestBonusMultiple(t *testing.T) {
	var test = "1d20 + 7 - 3"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '1d20[d1]+7-3= d2' d2 = [5-24]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}}, []int{7,-3}, sliceSum())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test tagged roll
// Minimum value 1, Maximum value 20
func TestTagBasic(t *testing.T) {
	var test = "1d20 Initiative"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: 'Initiative: 1d20[d1]= d1' d1 = [1-20]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}}, []int{}, sliceSum())
	if !strings.Contains(result, "Initiative") {
		t.Error(fmt.Sprintf("ERROR :: Tag not found : Expected: 'Initiative: 1d20[x]= x', Got: '%s'", result))
	}
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test tagged roll with bonus
// Minimum value 8, Maximum value 27
func TestTagMultiple(t *testing.T) {
	var test = "1d20 +7 Initiative"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: 'Initiative: 1d20[d1]+7= d1' d1 = [8-27]")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1,20}}, []int{7}, sliceSum())
	if !strings.Contains(result, "Initiative") {
		t.Error(fmt.Sprintf("ERROR :: Tag not found : Expected: 'Initiative: 1d20[x]+7= y', Got: '%s'", result))
	}
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test basic advantage roll
// Minimum value 1, Maximum value 20
func TestAdvantageBasic(t *testing.T) {
	var test = "m2d20"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d20[d1 d2]= d3' d3 = max(d1, d2)")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,20}}, []int{}, sliceTop())
	t.Log(fmt.Sprintf("Result: %s", result))
}

// Test basic advantage roll
// Minimum value 8, Maximum value 27
func TestAdvantageBonus(t *testing.T) {
	var test = "m2d20 + 7"
	t.Log(fmt.Sprintf("Test roll: %s", test))
	t.Log("Expected roll: '2d20[d1 d2]+7= d3' d3 = max(d1, d2)+7")
	var result, roller = Resolve(test)
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2,20}}, []int{7}, sliceTop())
	t.Log(fmt.Sprintf("Result: %s", result))
}
