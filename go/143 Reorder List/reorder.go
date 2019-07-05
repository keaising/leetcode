package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	walker, runner := head, head
	first := head
	for runner != nil && runner.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
	}
	cur := walker.Next
	walker.Next = nil

	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	for prev != nil && first != nil {
		p1, p2 := first.Next, prev.Next
		first.Next = prev
		prev.Next = p1
		first = p1
		prev = p2
	}
}
