package pearls

import "testing"

// Max sub array is [2:7] = 187
// var testArray = []int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}
var testTable = []struct {
	input    []int
	expected int
}{
	{[]int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}, 187},
	{[]int{-4, 4}, 4},
	{[]int{3, -3}, 3},
	{[]int{3, 3, 1}, 7},
	// This implementation considers an empty subarray to be a valid answer
	// so an input with all nagative entries returns 0. For certain domains
	// it's possible that the empty subarray isn't a valid answer.
	{[]int{-3, -3, -1}, 0},
}

func TestMaxSubArrayCubed(t *testing.T) {
	for _, testCase := range testTable {
		got := MaxSubArrayCubed(testCase.input)
		if got != testCase.expected {
			t.Errorf("Expected max sub array sum to be %v, got %v",
				testCase.expected, got)
		}
	}
}

func TestMaxSubArraySquared(t *testing.T) {
	for _, testCase := range testTable {
		got := MaxSubArraySquared(testCase.input)
		if got != testCase.expected {
			t.Errorf("Expected max sub array sum to be %v, got %v",
				testCase.expected, got)
		}
	}
}

func TestMaxSubArrayLinear(t *testing.T) {
	for _, testCase := range testTable {
		got := MaxSubArrayLinear(testCase.input)
		if got != testCase.expected {
			t.Errorf("Expected max sub array sum to be %v, got %v",
				testCase.expected, got)
		}
	}
}
