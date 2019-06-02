package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result = result ^ v
		fmt.Println(result, v)
	}
	return result
}

func main() {
	test := []int{1,2,1,3,3}
	singleNumber(test)
}
