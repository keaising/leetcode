package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	index := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	node := &TreeNode{
		Val:preorder[0],
		Left:buildTree(preorder[1:], inorder[:index]),
		Right:buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return node
}