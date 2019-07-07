package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{0, head}
	prev, cur := dummyHead.Next, dummyHead.Next.Next

	for cur != nil {
		// 要考虑相等的情况
		if cur.Val >= prev.Val {
			prev = cur
			cur = cur.Next
		} else {    // 需要在前面找位置进行插入
			// 暂存cur，待会用于插入。
			temp := cur
			// 将cur从链表中移除
			cur = cur.Next
			prev.Next = cur

			// 寻找插入的位置。
			prev1, cur1 := dummyHead, dummyHead.Next
			for cur1 != nil {
				if temp.Val < cur1.Val {    // 找到插入的位置，插入到cur1的前面
					prev1.Next = temp
					temp.Next = cur1
					break
				} else {
					prev1 = cur1
					cur1 = cur1.Next
				}
			}
		}
	}

	return dummyHead.Next
}