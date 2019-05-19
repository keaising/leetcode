package main

import "testing"

func TestSearch(t *testing.T) {
	inputs := []struct {
		nums   []int
		target int
		result int
	}{
		{
			nums:   []int{3, 1},
			target: 1,
			result: 1,
		},
		{
			nums:   []int{1, 3},
			target: 0,
			result: -1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			result: -1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			result: 0,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			result: 4,
		},
	}
	for _, input := range inputs {
		act := search(input.nums, input.target)
		if act != input.result {
			t.Error(input.nums, input.result, act)
		}
	}
}
