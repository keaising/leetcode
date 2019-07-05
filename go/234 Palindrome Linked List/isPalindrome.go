package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	s := []int{}
	p1, p2 := head, head
	for p2.Next != nil {
		s = append(s, p1.Val)
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	p1 = p1.Next
	for p1 != nil {
		if p1.Val != s[len(s)-1] {
			return false
		}
		s = s[:len(s)-1]
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	if p2 != nil {
		p1 = p1.Next
	}

	var prev *ListNode
	for p1 != nil {
		next := p1.Next
		p1.Next = prev
		prev = p1
		p1 = next
	}
	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	return true
}