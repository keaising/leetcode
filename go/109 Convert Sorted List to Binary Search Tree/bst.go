package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return sortArrayToBST(arr)
}

func sortArrayToBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   arr[len(arr)/2],
		Left:  sortArrayToBST(arr[:len(arr)/2]),
		Right: sortArrayToBST(arr[len(arr)/2+1:]),
	}
}
