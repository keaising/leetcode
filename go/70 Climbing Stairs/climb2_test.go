package main

import "testing"

func TestClimb(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, test := range tests {
		if test.output != climbStairs(test.input) {
			t.Error(test.input, climbStairs(test.input), test.output)
		}
	}
}
