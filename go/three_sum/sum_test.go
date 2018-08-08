package threesum

import (
		"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
	}{
		//{[]int{3, 2, 4}, 6},
		{[]int{-1,0,1,2,-1,-4}, 0},
	}
	for _, test := range tests {
		for k, v := range threeSum(test.input, test.target) {
			t.Log(k, v)
		}
	}
}
