package main

import "strconv"

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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	ans := new([]string)
	helper(root, ans, "")
	return *ans
}

func helper(root *TreeNode, ans *[]string, curr string) {
	if root == nil {
		return
	}
	if curr == "" {
		curr = strconv.Itoa(root.Val)
	} else {
		curr = curr + "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, curr)
	} else {
		helper(root.Left, ans, curr)
		helper(root.Right, ans, curr)
	}
}
