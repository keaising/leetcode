package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head.Next
	fore := head
	for curr != nil {
		if curr.Val == fore.Val {
			fore.Next = curr.Next
			curr = curr.Next
		} else {
			fore = fore.Next
			curr = curr.Next
		}
	}
	return head
}