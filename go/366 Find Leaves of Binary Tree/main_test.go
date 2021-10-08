package main

import (
	"reflect"
	"testing"
)

func Test_findLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: [][]int{
				{4, 5, 3},
				{2},
				{1},
			},
		},
		{
			name: "2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: [][]int{
				{2, 3},
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeaves(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
