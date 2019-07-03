package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	walker, runner, prev := head, head, head
	for runner != nil && runner.Next != nil {
		prev = walker
		walker = walker.Next
		runner = runner.Next.Next
	}
	if runner != nil {
		prev = walker
		walker = walker.Next
	}
	next := walker.Next
	for walker != nil {
		next = walker.Next
		walker.Next = prev
		prev = walker
		walker = next
	}
}
