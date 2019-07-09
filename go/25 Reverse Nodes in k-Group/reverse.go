package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	n := 0
	for i := head; i != nil; i = i.Next {
		n++
	}

	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	cur := prev.Next
	for ; n >= k; n = n - k {
		cur = prev.Next
		for m := 1; m < k; m++ {
			t := cur.Next
			cur.Next = t.Next
			t.Next = prev.Next
			prev.Next = t
		}

		prev = cur
	}
	return dummy.Next
}
