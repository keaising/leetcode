package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect bool
		no     int
	}{
		{
			nums:   []int{3, 1, 2, 2, 2},
			target: 1,
			expect: true,
			no:     12,
		},
		{
			nums:   []int{1, 3, 1, 1, 1},
			target: 3,
			expect: false,
			no:     11,
		},
		{
			nums:   []int{1, 3, 1, 1, 1},
			target: 2,
			expect: false,
			no:     10,
		},
		{
			nums:   []int{1, 3},
			target: 2,
			expect: true,
			no:     9,
		},
		{
			nums:   []int{1},
			target: 1,
			expect: true,
			no:     2,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			expect: false,
			no:     8,
		},
		{
			nums:   []int{5, 6, 7, 1, 2, 3, 4},
			target: 6,
			expect: true,
			no:     6,
		},
		{
			nums:   []int{5, 6, 7, 1, 2, 3, 4},
			target: 7,
			expect: true,
			no:     7,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			expect: true,
			no:     1,
		},
		{
			nums:   []int{1, 2},
			target: 1,
			expect: true,
			no:     3,
		},
		{
			nums:   []int{1, 2},
			target: 2,
			expect: true,
			no:     4,
		},
		{
			nums:   []int{1, 2, 3},
			target: 3,
			expect: true,
			no:     5,
		},
	}

	for _, test := range tests {
		result := search(test.nums, test.target)
		if result != test.expect {
			t.Error(test.no)
		}
	}
}
