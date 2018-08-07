package long

import "testing"

func TestLong(t *testing.T) {
	arr := []struct {
		Str string
		Len int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, a := range arr {
		act := lengthOfLongestSubstring(a.Str)
		if act != a.Len {
			t.Errorf("%s expect %d, actual %d", a.Str, a.Len, act)
		}
	}
}
