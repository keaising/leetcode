package median

import "testing"

func TestMedian(t *testing.T) {
	tests := []struct{
		a []int
		b []int
		expect float64
	} {
		{[]int{2}, []int{1,3},2},
		{[]int{2,4}, []int{1, 5},3},
		{[]int{2}, []int{5,6,7},5.5},
	}

	for _, test := range tests {
		act:= findMedianSortedArrays(test.a, test.b)
		if act != test.expect {
			t.Errorf("%d %d expect: %f, actual: %f",test.a, test.b, test.expect, act)
		}
	}
}
