package main

import "testing"

func TestValid(t *testing.T) {
	root := &TreeNode{
		Val: 13,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 14,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	if !isValidBST(root) {
		t.Error("Error!!!")
	}
}
