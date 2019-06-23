package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	temp := cur.Next
	for i := m; i < n; i++ {
		temp = cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return dummy.Next
}
