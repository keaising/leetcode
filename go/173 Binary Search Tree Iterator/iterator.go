package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	head  *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return BSTIterator{
		head:  root,
		stack: stack,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for this.head != nil {
		this.stack = append(this.stack, this.head)
		this.head = this.head.Left
	}
	node := this.stack[len(this.stack)-1]
	this.head = node.Right
	this.stack = this.stack[:len(this.stack)-1]
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.head != nil || len(this.stack) > 0
}
