package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil || root.Right != nil {
		return 1+bigger(maxDepth(root.Left), maxDepth(root.Right))
	}
	return 1
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}