package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root != nil {
		return [][]int{}
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
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	ret := levelOrder(tree)
	fmt.Println(ret)
}
