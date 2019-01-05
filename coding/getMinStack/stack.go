package stack

import (
	"github.com/golang-collections/collections/stack"
)

var (
	data = stack.New()
	min  = stack.New()
)

func push(i int) {
	data.Push(i)
	if getMin() >= i {
		min.Push(i)
	}
}

func getMin() int {
	return 0
}
