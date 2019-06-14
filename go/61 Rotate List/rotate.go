package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if k == 0 {
		return head
	}
	length := 1
	for curr := head; curr.Next != nil; curr = curr.Next {
		length++
	}
	breakPoint := length - k%length

	br := head
	last := head
	for i := 1; i < length; i++ {
		if i == breakPoint {
			br = last
		}
		last = last.Next
	}
	newHead := br.Next
	last.Next = head
	br.Next = nil

	return newHead

}

type ListNode struct {
	Val  int
	Next *ListNode
}
