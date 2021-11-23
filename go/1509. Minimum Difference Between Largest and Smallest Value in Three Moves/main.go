package main

import "sort"

func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	res := 0
	k := 3
	for i := 0; i < k; i++ {
		if res < nums[n-1-(k-i)]-nums[i] {
			res = nums[n-1-(k-i)] - nums[i]
		}
	}
	return res
}
