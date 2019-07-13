package main

import "testing"

func TestMax(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]byte
		expect int
	}{
		{
			name: "first",
			input: [][]byte{
				{'1', '0', '1', '1', '0', '1'},
				{'1', '1', '1', '1', '1', '1'},
				{'0', '1', '1', '0', '1', '1'},
				{'1', '1', '1', '0', '1', '0'},
				{'0', '1', '1', '1', '1', '1'},
				{'1', '1', '0', '1', '1', '1'},
			},
			expect: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expect != maximalSquare(tc.input) {
				t.Error("!!!!")
			}
		})
	}
}
