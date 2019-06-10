package foursum

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	ret := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	for k, v := range ret {
		t.Log(k, v)
	}
}

func TestNums(t *testing.T) {
	for _, v1 := range [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
	} {
		for k2, v2 := range v1 {
			t.Logf("%d: %d", k2, v2)
		}
	}
}
