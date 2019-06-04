package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	index := -1
	for i := range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			index = i
			break
		}
	}

	leftIn := inorder[:index]
	leftPost := postorder[:len(leftIn)]

	rightIn := inorder[index+1:]
	rightPost := postorder[len(leftPost) : len(postorder)-1]
	node := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(leftIn, leftPost),
		Right: buildTree(rightIn, rightPost),
	}
	return node
}
