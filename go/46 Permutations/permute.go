package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1,0}))
}

func permute(nums []int) [][]int {
	result = [][]int{}
	if len(nums) == 1 {
		return [][]int{
			nums,
		}
	}
	array = &nums
	generate(len(nums))
	return result
}

var result [][]int

var array *[]int

func generate(k int) {
	if k == 1 {
		ret := make([]int, len(*array))
		copy(ret, *array)
		result = append(result, ret)
	} else {
		generate(k - 1)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				swap(i, k-1)
			} else {
				swap(0, k-1)
			}
			generate(k - 1)
		}
	}
}

func swap(i, j int) {
	tem := (*array)[i]
	(*array)[i] = (*array)[j]
	(*array)[j] = tem
}
