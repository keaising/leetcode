package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	cur := head.Next.Next
	n1, n2 := p1, p2
	for i := 0; cur != nil; i++ {
		if i%2 == 0 {
			n1.Next = cur
			n1 = n1.Next
		} else {
			n2.Next = cur
			n2 = n2.Next
		}
		cur = cur.Next
	}
	n1.Next = p2
	n2.Next = nil

	return p1
}
