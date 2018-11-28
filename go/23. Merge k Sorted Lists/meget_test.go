package merge

import "testing"

func TestMerge(t *testing.T) {
	l1 := &ListNode{Val:2, Next:&ListNode{Val:5, Next:&ListNode{Val:7, Next:&ListNode{Val:10}}}}
	l2 := &ListNode{Val:1, Next:&ListNode{Val:3, Next:&ListNode{Val:4}}}
	l3 := &ListNode{Val:4, Next:&ListNode{Val:6, Next:&ListNode{Val:9}}}
	source := []*ListNode{l1, l2, l3}
	ret := mergeKLists(source)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

