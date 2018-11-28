package merge

import "testing"

func TestMerge(t *testing.T) {
	l1 := &ListNode{Val:2, Next:&ListNode{Val:5, Next:&ListNode{Val:7, Next:&ListNode{Val:10}}}}
	l2 := &ListNode{Val:1, Next:&ListNode{Val:3, Next:&ListNode{Val:4}}}
	ret := mergeTwoLists(l1, l2)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}
