package LongestPalindromic

import "testing"

func Test_Longest(t *testing.T) {
	tests := []struct {
		s   string
		max string
	}{
		{"dabax", "aba"},
	}
	for _, test := range tests {
		if ret := longestPalindrome(test.s); ret != test.max {
			t.Errorf("origin string: %s, want: %s, actual: %s", test.s, test.max, ret)
		}
	}
}
