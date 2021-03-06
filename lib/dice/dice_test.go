package dice

import (
	"sort"
	"strings"
	"testing"
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
func sliceHigher(dice int) func([]int) int {
	return func(slice []int) int {
		sort.Ints(slice)
		sum := sliceSum()
		return sum(slice[len(slice)-dice:])
	}
}

// Returns the higher value among the slice items
func sliceLower() func([]int) int {
	return func(slice []int) (lower int) {
		for _, item := range slice {
			if lower > item || lower == 0 {
				lower = item
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
			t.Errorf("ERROR :: Wrong number of dices : Expected: %d, Got: %d", dice, item.dice)
			// Checks the die faces
		} else if item.faces != faces {
			t.Errorf("ERROR :: Wrong number of die faces : Expected: %d, Got: %d", faces, item.faces)
			// Checks the value of the check result is lesser than the maximum possible
		} else if checkSum > (dice * faces) {
			t.Errorf(
				"ERROR :: Result outbounds dice maximum value : Expected: %d, Got: %d",
				dice*faces,
				checkSum,
			)
		}
		checkTotal += checkSum
	}
	// Checks that check total coincides with check results sum
	if checkTotal+sliceSum()(bonus) != r.total {
		t.Errorf("ERROR :: Wrong check total : Expected: %d, Got: %d", checkTotal, r.total)
	}
}

// --- Tests ---

// Test 1d20 roll
// Minimum value: 1, Maximum value: 20
func TestRollBasic(t *testing.T) {
	var test = "1d20"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '1d20[d1]= d1' d1 = [1-20]")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test 2d20+1d6 roll
// Minimum value: 3, Maximum value: 46
func TestRollMultiple(t *testing.T) {
	var test = "2d20+1d6"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]+1d6[d3]= d4' (d4=d1+d2+d3)")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}, {1, 6}}, []int{}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test default roll with no die
// Minimum value: 1, Maximum value: 20
func TestRollDefault(t *testing.T) {
	var test = ""
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '1d20[d1]= d1' d1 = [1-20]")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test 2d20+1d10+1d6 roll
// Minimum value: 4, Maximum value: 56
func TestRollWhitespaces(t *testing.T) {
	var test = "2d20 + 1d10 +1d6"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]+1d10[d3]+1d6[d4]= d5' (d5=d1+d2+d3+d4)")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}, {1, 10}, {1, 6}}, []int{}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test flat +5 bonus
// Minimum value: 6, Maximum value: 25
func TestBonusBasic(t *testing.T) {
	var test = "1d20+5"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '1d20[d1]+5= d2' d2 = [6-25]")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{5}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test flat -3 bonus
// Minimum value -1, Maximum value 17
func TestBonusNegative(t *testing.T) {
	var test = "2d10-3"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d10[d1]-3= d2' d2 = [-1-17]")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 10}}, []int{-3}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test multiple bonus
// Minimum value 5, Maximum value 24
func TestBonusMultiple(t *testing.T) {
	var test = "1d20 + 7 - 3"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '1d20[d1]+7-3= d2' d2 = [5-24]")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{7, -3}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test default bonus with no die
// Minimum value 5, Maximum value 24
func TestBonusDefault(t *testing.T) {
	var test = "+4"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '1d20[d1]+4= d2' d2 = [5-24]")
	var roller = resolve(test, "1d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{4}, sliceSum())
	t.Logf("Result: %s", roller.GetReply())
}

// Test tagged roll
// Minimum value 1, Maximum value 20
func TestTagBasic(t *testing.T) {
	var test = "1d20 Initiative"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: 'Initiative: 1d20[d1]= d1' d1 = [1-20]")
	var roller = resolve(test, "1d20")
	var result = roller.GetReply()
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{}, sliceSum())
	if !strings.Contains(result, "Initiative") {
		t.Errorf("ERROR :: Tag not found : Expected: 'Initiative: 1d20[x]= x', Got: '%s'", result)
	}
	t.Logf("Result: %s", result)
}

// Test tagged roll with bonus
// Minimum value 8, Maximum value 27
func TestTagMultiple(t *testing.T) {
	var test = "1d20 +7 Initiative"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: 'Initiative: 1d20[d1]+7= d2' d2 = [8-27]")
	var roller = resolve(test, "1d20")
	var result = roller.GetReply()
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{7}, sliceSum())
	if !strings.Contains(result, "Initiative") {
		t.Errorf("ERROR :: Tag not found : Expected: 'Initiative: 1d20[x]+7= y', Got: '%s'", result)
	}
	t.Logf("Result: %s", result)
}

// Test tagged roll with no die
// Minimum value 1, Maximum value 20
func TestTagDefalut(t *testing.T) {
	var test = "Initiative"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: 'Initiative: 1d20[d1]= d1' d1 = [1-20]")
	var roller = resolve(test, "1d20")
	var result = roller.GetReply()
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{1, 20}}, []int{}, sliceSum())
	if !strings.Contains(result, "Initiative") {
		t.Errorf("ERROR :: Tag not found : Expected: 'Initiative: 1d20[x]= x', Got: '%s'", result)
	}
	t.Logf("Result: %s", result)
}

// Test basic advantage roll
// Minimum value 1, Maximum value 20
func TestAdvantageBasic(t *testing.T) {
	var test = "h2d20"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]= d3' d3 = max(d1, d2)")
	var roller = resolve(test, "h2d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}}, []int{}, sliceHigher(1))
	t.Logf("Result: %s", roller.GetReply())
}

// Test complex advantage roll
// Minimum value 8, Maximum value 27
func TestAdvantageBonus(t *testing.T) {
	var test = "h2d20 + 7"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]+7= d3' d3 = max(d1, d2)+7")
	var roller = resolve(test, "h2d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}}, []int{7}, sliceHigher(1))
	t.Logf("Result: %s", roller.GetReply())
}

// Test default advantage roll with no die
// Minimum value 1, Maximum value 20
func TestAdvantageDefault(t *testing.T) {
	var test = ""
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]= d3' d3 = max(d1, d2)")
	var roller = resolve(test, "h2d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}}, []int{}, sliceHigher(1))
	t.Logf("Result: %s", roller.GetReply())
}

// Test complex advantage roll with multiple higher rolls
// Minimum value 3, Maximum value 18
func TestAdvantageComplex(t *testing.T) {
	var test = "3h4d6"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '4d6[d1 d2 d3 d4]= d5' d5 = 3bestOf(d1, d2, d3, d4)")
	var roller = resolve(test, "3h4d6")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{4, 6}}, []int{}, sliceHigher(3))
	t.Logf("Result: %s", roller.GetReply())
}

// Test basic disadvantage roll
// Minimum value 1, Maximum value 20
func TestDisadvantageBasic(t *testing.T) {
	var test = "l2d20"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]= d3' d3 = min(d1, d2)")
	var roller = resolve(test, "l2d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}}, []int{}, sliceLower())
	t.Logf("Result: %s", roller.GetReply())
}

// Test complex disadvantage roll
// Minimum value 7, Maximum value 26
func TestDisadvantageBonus(t *testing.T) {
	var test = "l2d20 +6"
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]+6= d3' d3 = min(d1, d2)+6")
	var roller = resolve(test, "l2d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}}, []int{6}, sliceLower())
	t.Logf("Result: %s", roller.GetReply())
}

// Test default disadvantage with no die
// Minimum value 1, Maximum value 20
func TestDisadvantageDefault(t *testing.T) {
	var test = ""
	t.Logf("Test roll: %s", test)
	t.Log("Expected roll: '2d20[d1 d2]= d3' d3 = min(d1, d2)")
	var roller = resolve(test, "l2d20")
	// Sends roll to checker
	checkRoll(t, roller, diceMatrix{{2, 20}}, []int{}, sliceLower())
	t.Logf("Result: %s", roller.GetReply())
}
