package structure

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	s := `1
2,3
4,nil,6,nil`
	root := BuildTree(s)
	arr := []*TreeNode{}
	var ret []int
	for root != nil {
		ret = append(ret, root.Val)
		if root.Right != nil {
			arr = append(arr, root.Right)
		}
		root = root.Left
		if root == nil && len(arr) > 0 {
			root = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		}
	}
	t.Log(ret)
}

func TestStr2Node(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output *TreeNode
	}{
		{
			"1st",
			"1",
			&TreeNode{
				Val: 1,
			},
		},
		{
			"2nd",
			"0",
			&TreeNode{
				Val: 0,
			},
		},
		{
			"3rd",
			"nil",
			nil,
		},
		{
			"4th",
			"4",
			&TreeNode{
				Val: 4,
			},
		},
		{
			"5th",
			"nil",
			&TreeNode{
				Val: 4,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := str2Node(tc.input)
			if output != nil && tc.output != nil {
				if output.Val != tc.output.Val {
					t.Errorf("not equal! input:%q, output: %v, expect: %v", tc.input,
						output.Val, tc.output.Val)
				}
			} else {
				if output == nil && tc.output == nil {

				} else {
					t.Errorf("not all nil! input:%q, output: %t, expect: %t", tc.input,
						output == nil, tc.output == nil)
				}
			}
		})
	}
}
