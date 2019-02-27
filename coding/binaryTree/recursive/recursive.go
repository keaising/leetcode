package recursive

import (
	"../Data"
	"fmt"
)

func preOrderCur(head *Data.Node) {
	if head == nil {
		return
	}
	fmt.Println(head.Val, " ")
	preOrderCur(head.Left)
	preOrderCur(head.Right)
}

func inOrderCur(head *Data.Node) {
	if head == nil {
		return
	}
	inOrderCur(head.Left)
	fmt.Println(head.Val, " ")
	inOrderCur(head.Right)
}

func postOrderCur(head *Data.Node) {
	if head == nil {
		return
	}
	postOrderCur(head.Left)
	postOrderCur(head.Right)
	fmt.Println(head.Val, " ")
}
