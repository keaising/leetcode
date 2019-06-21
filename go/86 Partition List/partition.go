package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	less, bigger := new(ListNode), new(ListNode)
	l, b := less, bigger
	for head != nil {
		if head.Val < x {
			less.Next = &ListNode{
				Val: head.Val,
			}
			less = less.Next
		} else {
			bigger.Next = &ListNode{
				Val: head.Val,
			}
			bigger = bigger.Next
		}
		head = head.Next
	}
	less.Next = b.Next
	head = l.Next
	return head
}
