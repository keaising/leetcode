package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	one:=head
	two:=head
	for one != nil && two != nil {
		one = one.Next
		two = two.Next.Next
		if one == two {
			return true
		}
	}
	return false
}
