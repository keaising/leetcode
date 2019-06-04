package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	node := &TreeNode{
		Val:nums[len(nums)/2],
	}
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2+1:]
	node.Left = sortedArrayToBST(left)
	node.Right = sortedArrayToBST(right)
	return node
}