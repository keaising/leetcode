package main

type MyStack struct {
	s []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	s := MyStack{}
	return s
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.s = append(this.s, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.s) == 0 {
		return 0
	} else {
		item := this.s[len(this.s)-1]
		this.s = this.s[:len(this.s)-1]
		return item
	}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.s[len(this.s)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.s) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
