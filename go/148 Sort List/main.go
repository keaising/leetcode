package main

import "fmt"

func main() {
	list := buildListNode([]int{1, 2, 3, 4})
	list = reverse(list)
	fmt.Println(buildNums(list))
}

func reverse(head *ListNode) *ListNode {
	prev := &ListNode{
		Next:head,
	}
	cur := head
	for cur != nil && cur.Next != nil {
		t := cur.Next
		cur.Next = t.Next
		t.Next = prev.Next
		prev.Next = t
	}
	return prev.Next
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
