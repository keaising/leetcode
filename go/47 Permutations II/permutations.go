package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	helper([]int{}, nums, &ans)
	return ans
}

func helper(cur []int, nums []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, cur)
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && (nums)[i] == (nums)[i-1] {
			continue
		}
		helper(append(append([]int{}, cur...), (nums)[i]),
			append(append([]int{}, nums[:i]...), nums[i+1:]...),
			ans)
	}
}
