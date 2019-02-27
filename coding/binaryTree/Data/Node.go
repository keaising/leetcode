package Data

type Node struct {
	Val   int32
	Left  *Node
	Right *Node
}

var NormalTree *Node = &Node{
	Val: 1,
	Left: &Node{
		Val: 2,
		Left: &Node{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &Node{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &Node{
		Val: 3,
		Left: &Node{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &Node{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}
