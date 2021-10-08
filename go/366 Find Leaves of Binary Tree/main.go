package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var current []int
	for root != nil {
		if isLeaf(root, &current) {
			root = nil
		}
		result = append(result, current)
		current = []int{}
	}
	return result
}

func isLeaf(node *TreeNode, current *[]int) bool {
	if node == nil {
		return true
	}
	if node.Left == nil && node.Right == nil {
		*current = append(*current, node.Val)
		log.Println("is leaf", node.Val, "current", current)
		return true
	}
	if isLeaf(node.Left, current) {
		node.Left = nil
	}
	if isLeaf(node.Right, current) {
		node.Right = nil
	}
	return false
}
