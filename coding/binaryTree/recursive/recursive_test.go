package recursive

import (
	"../Data"
	"testing"
)

func TestPreOrderCur(t *testing.T) {
	preOrderCur(Data.NormalTree)
}

func TestInOrderCur(t *testing.T) {
	inOrderCur(Data.NormalTree)
}

func TestPostOrderCur(t *testing.T) {
	postOrderCur(Data.NormalTree)
}
