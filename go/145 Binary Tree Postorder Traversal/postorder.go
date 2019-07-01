package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ans)
	helper(root.Right, ans)
	*ans = append(*ans, root.Val)
}

