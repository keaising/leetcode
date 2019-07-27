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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	ary := make([]int, n)
	for i := 0; i < n; i++ {
		ary[i] = i + 1;
	}

	return helper(ary)
}

func helper(ary []int) []*TreeNode {
	if len(ary) == 0 {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}
	for i := 0; i < len(ary); i++ {
		for _, left := range helper(ary[:i]) {
			for _, right := range helper(ary[i+1:]) {
				root := &TreeNode{
					Val:   ary[i],
					Left:  left,
					Right: right,
				}

				result = append(result, root)
			}
		}
	}

	return result
}
