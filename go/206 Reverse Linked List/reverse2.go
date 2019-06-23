package main

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	array := [](*ListNode){}
	for head != nil {
		array = append(array, head)
		head = head.Next
	}
	head = &ListNode{
		Val:array[len(array)-1].Val,
	}
	curr := head
	for i:= len(array)-2; i>=0;i--{
		curr.Next = &ListNode{
			Val:array[i].Val,
		}
		curr = curr.Next
	}
	return head
}
