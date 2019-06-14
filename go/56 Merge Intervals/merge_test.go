package main

import (
	"github.com/keaising/leetcode/go/structure"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expect    [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expect:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expect:    [][]int{{1, 5}},
		},
	}

	for _, test := range tests {
		if structure.TwoDEqual(test.expect, merge(test.intervals)) {

		}
	}
}
