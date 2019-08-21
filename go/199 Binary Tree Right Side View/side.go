package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	helper(root, &res, 0)
	return res
}

func helper(root *TreeNode, res *[]int, depth int) {
	if root == nil {
		return
	}
	if depth == len(*res) {
		*res = append(*res, root.Val)
	}
	helper(root.Right, res, depth+1)
	helper(root.Left, res, depth+1)
}
