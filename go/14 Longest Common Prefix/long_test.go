package long

import "testing"

func TestLong(t *testing.T) {
	strs := []string {"leets", "leetcode", "leet", "leeds"}
	actual := longestCommonPrefix(strs)
	if actual != "lee" {
		t.Errorf("wrong! want:%s, actual:%s", "lee", actual)
	}
}
