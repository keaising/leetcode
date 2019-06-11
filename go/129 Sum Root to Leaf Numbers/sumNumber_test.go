package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		node   *TreeNode
		expect int
	}{
		{
			node: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right:&TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			expect: 25,
		},
	}

	for _, test := range tests {
		if test.expect != sumNumbers(test.node) {
			t.Error(test.expect)
		}
	}

}
