package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	cur, prev, post := head, head, head.Next
	for prev != nil && prev.Val == val {
		prev = prev.Next
	}
	if prev == nil {
		return prev
	}
	head = prev
	cur = prev.Next
	for cur != nil {
		post = cur.Next
		if cur.Val == val {
			prev.Next = post
		} else {
			prev = cur
		}
		cur = post
	}
	return head
}
