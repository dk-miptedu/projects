package main

import "testing"

func TestMax(t *testing.T) {
	testCases := []struct {
		x, y, expected int
	}{
		{100000000000000000, 2, 100000000000000000},
		{1e10, 2e10, 2e10},
		{3e10, 2e10, 3e10},
		{20, 19, 20},
		{19, 20, 20},
		{-1, -10, -1},
		{-10, -1, -1},
		{-10, -10, -10},
		{0, 0, 0},
		{100, 100, 100},
	}

	for somelse, testCase := range testCases {
		result := Max(testCase.x, testCase.y)
		if result != testCase.expected {
			t.Errorf("Max(%d, %d) = %d; expected %d", testCase.x, testCase.y, result, testCase.expected)
		} else {
			println(somelse, "Max(", testCase.x, ",", testCase.y, ") = ", result, "; expected ", testCase.expected)

		}
	}
}
