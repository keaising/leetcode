package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if q.Val != p.Val {
		return false
	}
	return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
}
