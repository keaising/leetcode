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
	q.fill(root)
	curr := root
	for _, node := range q {
		curr.Left = nil
		curr.Right = node
		curr = node
	}
}

func (q Queue) fill(root *TreeNode) Queue {
	if root != nil {
		q = q.in(root)
		q = q.fill(root.Left)
		q = q.fill(root.Right)
	}
	return q
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
