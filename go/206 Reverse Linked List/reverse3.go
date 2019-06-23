package main

func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev:= head
	cur := head.Next
	temp := head.Next.Next

	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	head.Next = nil
	return prev
}
