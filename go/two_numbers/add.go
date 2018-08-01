package add

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
	var res = new(ListNode)
	cur := res
	carry := 0
	next := func(p *ListNode) *ListNode {
		if p == nil {
			return nil
		}
		return p.Next
	}

	for l1 != nil || l2 != nil || carry > 0 {
		value := carry
		switch {
		case l1 != nil && l2 != nil:
			value += l1.Val + l2.Val
		case l1 != nil:
			value += l1.Val
		case l2 != nil:
			value += l2.Val
		}
		carry = value / 10
		cur.Val = value % 10
		l1 = next(l1)
		l2 = next(l2)

		if l1 != nil || l2 != nil || carry > 0 {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	return res
}
