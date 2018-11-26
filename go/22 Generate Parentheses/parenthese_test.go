package parenthes


import "testing"

func TestGenerateParenthesis(t *testing.T) {
	ret := generateParenthesis(1)
	t.Log(len(ret))
	t.Log(ret)
}
