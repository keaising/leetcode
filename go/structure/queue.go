package structure

type Queue []*TreeNode

func (q *Queue) In(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Pop() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}
