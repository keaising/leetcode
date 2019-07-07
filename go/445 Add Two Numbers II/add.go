package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := length(l1)
	n2 := length(l2)
	if n2 > n1 {
		l1, l2 = l2, l1
	}
	l1 = reverse(l1)
	l2 = reverse(l2)
	head := l1
	carry := 0
	for l2 != nil {
		l1.Val = l1.Val + l2.Val + carry
		if l1.Val >= 10 {
			l1.Val = l1.Val % 10
			carry = 1
		} else {
			carry = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if carry == 1 {
		for l1 != nil {
			l1.Val += carry
			if l1.Val > 9 {
				l1.Val %= 10
				carry = 1
			} else {
				carry = 0
			}
			l1 = l1.Next
		}
		if carry == 1 {
			curr := head
			for curr != nil && curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
		}
	}
	head = reverse(head)
	return head
}

func length(list *ListNode) int {
	n := 0
	for list != nil {
		list = list.Next
		n++
	}
	return n
}

func reverse(list *ListNode) *ListNode {
	var prev *ListNode
	next := list.Next
	for list != nil {
		next = list.Next
		list.Next = prev
		prev = list
		list = next
	}
	return prev
}
