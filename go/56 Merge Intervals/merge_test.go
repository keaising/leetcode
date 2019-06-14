package main

import (
	"github.com/keaising/leetcode/go/structure"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expect    [][]int
		no        int
	}{
	{
	intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		expect:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		no:        1,
	},
	{
	intervals: [][]int{{1, 4}, {4, 5}},
		expect:    [][]int{{1, 5}},
		no:        2,
	},
}

	for _, test := range tests {
		actual := merge(test.intervals)
		if !structure.TwoDEqual(test.expect, actual) {
			t.Error(test.no, actual)
		}
	}
}
