package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	total := 0
	sum(root, 0, &total)
	return total
}

func sum(root *TreeNode, curr int, total *int) {
	if root == nil {
		return
	}
	curr = curr*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*total += curr
		return
	}
	sum(root.Left, curr, total)
	sum(root.Right, curr, total)
}
