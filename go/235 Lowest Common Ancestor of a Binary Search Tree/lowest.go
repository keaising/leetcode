package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > bigger(p.Val, q.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < smaller(p.Val, q.Val) {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func smaller(a, b int) int {
	if a > b {
		return b
	}
	return a
}
