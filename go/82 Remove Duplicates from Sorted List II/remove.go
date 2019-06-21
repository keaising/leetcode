package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	slow := new(ListNode)
	slow.Next = head
	fast := slow
	for fast.Next != nil {
		if fast.Next.Next != nil && fast.Next.Val == fast.Next.Next.Val {
			value := fast.Next.Val
			for fast.Next != nil && fast.Next.Val == value {
				fast.Next = fast.Next.Next
			}
		} else {
			fast = fast.Next
		}
	}
	return slow.Next
}
