package substring

import "testing"

func TestSub(t *testing.T) {
	ret := findSubstring("barfootheegoodgood", []string{"thee", "good", "good"})
	t.Log(ret)
}
