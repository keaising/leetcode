package recursive

import "fmt"

func preOrderCur(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.Val, " ")
	preOrderCur(head.Left)
	preOrderCur(head.Right)
}

func inOrderCur(head *Node) {
	if head == nil {
		return
	}
	inOrderCur(head.Left)
	fmt.Println(head.Val, " ")
	inOrderCur(head.Right)
}

func postOrderCur(head *Node) {
	if head == nil {
		return
	}
	postOrderCur(head.Left)
	postOrderCur(head.Right)
	fmt.Println(head.Val, " ")
}
