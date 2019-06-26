package main

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, smaller(this.min[len(this.min)-1], x))
}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 0
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}
