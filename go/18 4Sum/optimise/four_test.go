package optimise

import (
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	for k, v := range fourSum([]int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5}, -13) {
		if v != nil {
			t.Log(k, v)
		}
	}
}

func TestInts2(t *testing.T) {
	ret := []int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5}
	sort.Ints(ret)
	t.Log(ret)
}
