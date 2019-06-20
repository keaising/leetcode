package main

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
