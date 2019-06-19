package main

import "testing"

func TestSubset(t *testing.T) {
	tests := []struct {
		nums   []int
		expect [][]int
		no     int
	}{
		{
			nums: []int{1, 2, 3},
			expect: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}

	for _, test := range tests {
		result := subsets(test.nums)
		t.Error(result)
	}
}
