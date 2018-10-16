package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make(map[int]*ListNode)

	for i := 1; head != nil; i++ {
		nodes[i] = head
		head = head.Next
	}

	nth := len(nodes) - n + 1
	if nth == 1 {
		return nodes[2]
	} else if len(nodes) == nth {
		nodes[nth].Next = nil
		return nodes[1]
	}
	nodes[nth-1].Next = nodes[nth].Next
	return nodes[1]
}
