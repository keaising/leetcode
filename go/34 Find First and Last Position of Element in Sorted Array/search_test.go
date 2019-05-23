package main

import "testing"

func TestSearchRange(t *testing.T) {
	inputs := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{1, 2, 3},
			target: 2,
			result: []int{1, 1},
		},
		{
			nums:   []int{1, 3},
			target: 1,
			result: []int{0, 0},
		},
		{
			nums:   []int{1, 3},
			target: 0,
			result: []int{-1, -1},
		},
		{
			nums:   []int{4, 4, 6, 7},
			target: 4,
			result: []int{0, 1},
		},
		{
			nums:   []int{1, 2, 3, 4, 4, 6},
			target: 4,
			result: []int{3, 4},
		},
		{
			nums:   []int{4, 4},
			target: 4,
			result: []int{0, 1},
		},
		{
			nums:   []int{1, 4},
			target: 4,
			result: []int{1, 1},
		},
	}
	for _, input := range inputs {
		act := searchRange(input.nums, input.target)
		if !equal(act, input.result) {
			t.Error(input.nums, input.target, input.result, act)
		}
	}
}

func equal(act, exp []int) bool {
	if act[0] == exp[0] && act[1] == exp[1] {
		return true
	}
	return false
}
