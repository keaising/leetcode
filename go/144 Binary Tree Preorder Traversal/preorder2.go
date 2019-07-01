package main

func preorderTraversal2(root *TreeNode) []int {
	a := []int{}
	r := [](*TreeNode){}
	for root != nil {
		a = append(a, root.Val)
		if root.Right != nil {
			r = append(r, root.Right)
		}
		root = root.Left
		if root == nil && len(r) > 0 {
			root = r[len(r)-1]
			r = r[:len(r)-1]
		}
	}
	return a
}
