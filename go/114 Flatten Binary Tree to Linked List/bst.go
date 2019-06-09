package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	q := Queue{}
	q = q.in(root.Left)
	q = q.in(root.Right)
	curr := root
	for len(q) > 0 {
		node, q := q.pop()
		if node != nil {
			q = q.in(node.Left)
			q = q.in(node.Right)
			curr.Left = nil
			curr.Right = node
			curr = node
		}
	}
}

func (q Queue) fill(root *TreeNode) Queue {
	if root != nil {
		q = q.in(root)
		
	}
}

type Queue []*TreeNode

func (q Queue) in(node *TreeNode) Queue {
	q = append(q, node)
	return q
}

func (q Queue) pop() (*TreeNode, Queue) {
	if len(q) == 0 {
		return nil, q
	}
	return q[0], q[1:]
}
