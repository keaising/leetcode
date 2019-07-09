package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{
		Next:head,
	}
	helper(&prev, nil)

	return prev.Next
}

func helper(left, right **ListNode) {
	for pivot, cur := left, left.Next; cur != nil; cur = cur.Next {
		if cur.Val > pivot.Val {

		}
	}
	return nil
}
