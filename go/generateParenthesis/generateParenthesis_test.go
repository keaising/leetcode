package generateParenthesis

import (
	"fmt"
	"testing"
)

func TestForNeibor(t *testing.T) {
	ret := generateParenthesis(4)
	fmt.Println(ret)
	fmt.Println(len(ret))
}
