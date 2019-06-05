package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := depthBalance(root)
	return ok
}

func depthBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	d1, ok1 := depthBalance(node.Left)
	d2, ok2 := depthBalance(node.Right)
	if ok1 != ok2 || abs(d1-d2) > 1 {
		return 0, false
	}
	return bigger(d1+1, d2+1), true
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
