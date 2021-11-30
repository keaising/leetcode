package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []*TreeNode
var cache map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res = []*TreeNode{}
	cache = make(map[string]int)

	_ = traverse(root)
	return res
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	left := traverse(root.Left)
	right := traverse(root.Right)

	node := left + fmt.Sprintf("%d", root.Val) + right
	cache[node] += 1
	if cache[node] == 2 {
		res = append(res, root)
	}
	return node
}
