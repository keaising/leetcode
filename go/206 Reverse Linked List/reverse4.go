package main

func reverseList4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, curr, next := head, head.Next, head.Next.Next
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next = nil
	return prev
}
