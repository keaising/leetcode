package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	arr := []*TreeNode{}
	arr = append(arr, root)
	for len(arr) > 0 {
		tempRet := []int{}
		tempArr := []*TreeNode{}
		for _, node := range arr {
			if node != nil {
				tempRet = append(tempRet, node.Val)
				if node.Left != nil {
					tempArr = append(tempArr, node.Left)
				}
				if node.Right != nil {
					tempArr = append(tempArr, node.Right)
				}
			}
		}
		arr = tempArr
		result = append(result, tempRet)
	}
	return reverse(result)
}

func reverse(arr [][]int) [][]int {
	ret := [][]int{}
	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i])
	}
	return ret
}