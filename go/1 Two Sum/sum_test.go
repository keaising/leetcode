package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect []int
	}{
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}
	for _, test := range tests {
		if act := twoSum(test.input, test.target); !reflect.DeepEqual(act, test.expect) {
			t.Fatalf("arr %v, want %v, actual %v", test.input, test.expect, act)
		}
	}
}
