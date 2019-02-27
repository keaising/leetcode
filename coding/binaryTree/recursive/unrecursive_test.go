package recursive

import (
	"../Data"
	"testing"
)

func TestPreOrderUnCur(t *testing.T) {
	preOrderUncur(Data.NormalTree)
}

func TestInOrderUnCur(t *testing.T) {
	inOrderUncur(Data.NormalTree)
}
