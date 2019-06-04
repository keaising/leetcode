package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	stack := Stack{}
	stack = append(stack, root)
	turn := true
	for len(stack) > 0 {
		curr := []int{}
		temp := []*TreeNode{}
		for len(stack) > 0 {
			s, node := stack.Pop()
			stack = s
			curr = append(curr, node.Val)
			if turn {
				if node.Left != nil {
					temp = append(temp, node.Left)
				}
				if node.Right != nil {
					temp = append(temp, node.Right)
				}
			} else {
				if node.Right != nil {
					temp = append(temp, node.Right)
				}
				if node.Left != nil {
					temp = append(temp, node.Left)
				}
			}
		}
		turn = !turn
		stack = temp
		ret = append(ret, curr)
	}
	return ret
}

type Stack []*TreeNode

func (s Stack) Push(node *TreeNode) Stack {
	s = append(s, node)
	return s
}

func (s Stack) Pop() (Stack, *TreeNode) {
	node := s[len(s)-1]
	s = s[:len(s)-1]
	return s, node
}
