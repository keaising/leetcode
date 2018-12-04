package str

import "testing"

func TestStr(t *testing.T) {
	ret := strStr("aaaaabcaa", "abc")
	t.Log(ret)
}
