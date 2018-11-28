/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tmp := new(ListNode)
	var start = tmp
	if l1.Val > l2.Val {
		tmp.Val = l2.Val
		l2 = l2.Next
 	} else {
 		tmp.Val = l1.Val
 		l1 = l1.Next
	}
	for {
		if l1 == nil {
			tmp.Next = l2
			break
		}
		if l2 == nil {
			tmp.Next = l1
			break
		}
		if l1.Val > l2.Val {
			tmp.Next = &ListNode{Val:l2.Val}
			tmp = tmp.Next
			l2 = l2.Next
		} else {
			tmp.Next = &ListNode{Val:l1.Val}
			tmp = tmp.Next
			l1 = l1.Next
		}
	}
	return start
}
