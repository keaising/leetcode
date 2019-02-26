package recursive

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func preOrderUncur(head *Node) {
	fmt.Println("pre-order: ")
	if head != nil {
		stack := stack.New()
		stack.Push(head)
		for stack.Len() > 0 {
			head, _ = stack.Pop().(*Node)
			fmt.Println(head.Val)
			if head.Right != nil {
				stack.Push(head.Right)
			}
			if head.Left != nil {
				stack.Push(head.Left)
			}
		}
	}
}

func inOrderUncur(head *Node) {
	fmt.Println("in-order: ")
	if head != nil {
		stack := stack.New()
		for stack.Len() > 0 || head != nil {
			if head != nil {
				stack.Push(head)
				head = head.Left
			} else {
				head = stack.Pop().(*Node)
				fmt.Println(head.Val)
				head = head.Right
			}
		}
	}
}
