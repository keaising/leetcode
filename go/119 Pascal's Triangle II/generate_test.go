package main

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		number int
		result []int
	}{
		{
			number: 0,
			result: []int{1},
		},
		{
			number: 1,
			result: []int{1, 1},
		},
		{
			number: 2,
			result: []int{1, 2, 1},
		},
		{
			number: 3,
			result: []int{1, 3, 3, 1},
		},
	}

	for _, test := range tests {
		if fmt.Sprintf("%q", test.result) != fmt.Sprintf("%q", getRow(test.number)) {
			t.Error(test.number, test.result, getRow(test.number))
		}
	}
}
