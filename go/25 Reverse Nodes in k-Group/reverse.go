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
	for i := head; i != nil; i = i.Next {
		n += 1
	}

	dummy := &ListNode{
		Next: head,
	}
	prev := dummy.Next
	cur := head.Next
	for ; n > k; n = n - k {
		for m := 0; m < k; m++ {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}
