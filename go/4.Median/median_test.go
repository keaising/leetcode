package median

import "testing"

func TestMedian(t *testing.T) {
	ret := findMedianSortedArrays([]int{2}, []int{1,3})
	if ret != 2.0 {
		t.Fatalf("expect:", ret)
	}
}
