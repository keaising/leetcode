package structure

import (
	"strconv"
	"strings"
)

func BuildTree(str string) *TreeNode {
	if len(str) == 0 {
		return nil
	}
	lines := strings.Split(str, "\n")
	var arr []*TreeNode
	root := str2Node(lines[0])
	arr = append(arr, root)
	for i := 1; i < len(lines); i++ {
		length := len(arr)
		nexts := str2Nodes(lines[i])
		for j := 0; j < length; j++ {
			if arr[j] != nil {
				arr[j].Left = nexts[2*j]
				arr[j].Right = nexts[2*j+1]
			}
		}
		arr = nexts
	}
	return root
}

func str2Nodes(s string) []*TreeNode {
	var ret []*TreeNode
	arr := strings.Split(s, ",")
	for _, a := range arr {
		ret = append(ret, str2Node(a))
	}
	return ret
}

func str2Node(s string) *TreeNode {
	if s == "nil" {
		return nil
	} else {
		i, _ := strconv.Atoi(s)
		return &TreeNode{
			Val:   i,
			Left:  nil,
			Right: nil,
		}
	}
}
