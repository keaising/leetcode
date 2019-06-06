package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	find(root, sum, []int{}, &result)
	return result
}

func find(node *TreeNode, sum int, arr []int, result *[][]int) {
	if node == nil {
		return
	}
	arr = append(arr, node.Val)
	if node.Val == sum && node.Left == nil && node.Right == nil {

		//*result = append(*result, append([]int(nil),arr...))
		*result = append(*result, append([]int{}, arr...))
		return
	}
	find(node.Left, sum-node.Val, append([]int{}, arr...), result)
	find(node.Right, sum-node.Val, append([]int{}, arr...), result)
}
