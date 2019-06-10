package structure

func buildTreeFromInts(arr []int){
	q := Queue{}
	root := &TreeNode{
		Val:arr[0],
		Left:nil,
		Right:nil,
	}
	q.In(root)
	i := 1
	for len(q) != 0 {
		node, q:= q.Pop()
		q = q
		if node != nil {

		}
	}
}
