package main

import "testing"

func TestIs(t *testing.T) {
	tests := []struct {
		str    string
		expect bool
	}{
		{
			str:    "",
			expect: true,
		},
		{
			str:    "1",
			expect: true,
		},
		{
			str:    "A man, a plan, a canal: Panama",
			expect: true,
		},
		{
			str:    "race a car",
			expect: false,
		},
		{
			str:    ".,",
			expect: true,
		},
	}

	for _, test := range tests {
		if test.expect != isPalindrome(test.str) {
			t.Error(test.str, test.expect, isPalindrome(test.str))
		}
	}
}
