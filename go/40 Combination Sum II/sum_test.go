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
				{7},
			},
			no: 1,
		},
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expect: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
			no: 2,
		},
	}
	for _, test := range tests {
		act := combinationSum2(test.candidates, test.target)
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
