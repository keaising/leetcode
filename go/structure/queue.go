package structure

type Queue []*TreeNode

func (q Queue) In(node *TreeNode) Queue {
	q = append(q, node)
	return q
}

func (q Queue) Pop() (*TreeNode, Queue) {
	if len(q) == 0 {
		return nil, q
	}
	return q[0], q[1:]
}

