package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := 0
	for head != nil {
		head = head.Next
		n++
	}

	times := n / k
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	cur := head
	for i := 0; i < times; i++ {
		for m := 0; m < k; m++ {
			next := cur.Next
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}
