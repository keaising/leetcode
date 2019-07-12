package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}