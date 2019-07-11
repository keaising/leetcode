package main

import "testing"

func TestLarge(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		ans  int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 3},
	}

	for _, test := range tests {
		if test.ans != findKthLargest(test.nums, test.k) {
			t.Error("!!!")
		}
	}
}
