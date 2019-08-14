package main

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	cnt := n - m
	for m-1 > 0 {
		pre = pre.Next
		m--
	}
	cur := pre.Next
	for cnt > 0 {
		pre.Next, cur.Next.Next, cur.Next = cur.Next, pre.Next, cur.Next.Next
		cnt--
	}
	return dummy.Next
}
