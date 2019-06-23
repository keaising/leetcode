package main

func inorderTraversal2(root *TreeNode) []int {
	ret := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, root.Val)
			root = root.Right
		}
	}
	return ret
}
