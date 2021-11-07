package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor([]int{1, 2, 3})
	param_1 := obj.Reset()
	fmt.Println(param_1)
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Shuffle())
}

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.origin
}

func (this *Solution) Shuffle() []int {
	var shuffle = make([]int, len(this.origin))
	copy(shuffle, this.origin)

	for i := len(shuffle) - 1; i > 0; i-- {
		r := rand.Intn(i + 1)
		shuffle[i], shuffle[r] = shuffle[r], shuffle[i]
	}
	return shuffle
}
