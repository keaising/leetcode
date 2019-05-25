package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expect     [][]int
		no         int
	}{
		{
			candidates: []int{2, 3, 7},
			target:     7,
			expect: [][]int{
				{2, 2, 3},
				{7},
			},
			no: 1,
		},
	}
	for _, test := range tests {
		act := combinationSum(test.candidates, test.target)
		if !equal2D(act, test.expect) {
			t.Error(test.no, act, test.expect)
		}
	}

}

func equal2D(act, expect [][]int) bool {
	if len(act) != len(expect) {
		return false
	}
	for i := range act {
		if !equal1D(act[i], expect[i]) {
			return false
		}
	}
	return true
}

func equal1D(act, expect []int) bool {
	if len(act) != len(expect) {
		return false
	}
	for i := range act {
		if act[i] != expect[i] {
			return false
		}
	}
	return true
}
