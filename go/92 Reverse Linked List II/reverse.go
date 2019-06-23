package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	head = dummy

	for i := 1; i < m; i++ {
		head = head.Next
	}

	pre := head.Next
	curr := pre.Next
	for i := 0; i < n-m; i++ {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	head.Next.Next = curr
	head.Next = pre
	head = dummy.Next
	return head
}
