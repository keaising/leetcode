package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	end := head
	for end.Next != nil {
		end = end.Next
	}
	return helper(head, head, end)
}

func helper(head *ListNode, left, right *ListNode) *ListNode {
	for pivot, cur := left, left.Next; cur != nil; cur = cur.Next {
		if cur.Val > pivot.Val {

		}
	}
}
