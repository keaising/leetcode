package main


type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != nil || l2 != nil {
		if l1 == l2 {
			return l1
		}
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}
		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}
	return nil
}