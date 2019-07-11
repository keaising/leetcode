package main

import "sort"

func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

