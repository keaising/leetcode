package optimise

import "testing"

func TestOptimis(t *testing.T) {
	arr := []struct {
		s string
		i int
	}{
		//{"qqwwerttyuioo", 5},
		//{"wwwqqq", 2},
		{"aw", 2},
	}
	for _, a := range arr {
		i := lengthOfLongestSubstring(a.s)
		if i != a.i {
			t.Errorf("%s max length expect %d, acutal: %d", a.s, a.i, i)
		}
	}
}
