package main

import (
	"testing"
)

func TestReorder(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			ans:  []int{1, 7, 2, 6, 3, 5, 4},
		},
	}
	for _, test := range tests {
		list := buildListNode(test.nums)
		reorderList(list)
		nums := buildNums(list)
		if sliceEqual(test.ans, nums) {
			t.Error("!!!!")
		}
	}
}

func buildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		node = node.Next
	}
	return head
}

func buildNums(head *ListNode) []int {
	if head == nil {
		return nil
	}
	ans := []int{}
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
