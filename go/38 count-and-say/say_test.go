package main

import "testing"

func TestCountAndSay(t *testing.T) {
	inputs := []struct {
		n      int
		expect string
	}{
		{
			n:      1,
			expect: "1",
		},
		{
			n:      2,
			expect: "11",
		},
		{
			n:      3,
			expect: "21",
		},
		{
			n:      4,
			expect: "1211",
		},
		{
			n:      5,
			expect: "111221",
		},
		{
			n:      6,
			expect: "312211",
		},
	}
	for _, input := range inputs {
		act := countAndSay(input.n)
		if act != input.expect {
			t.Error(input.n, input.expect, act)
		}
	}
}
