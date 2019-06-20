package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	n, repeated := 1, -9999999
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if repeated != nums[i] {
				nums[n] = nums[i]
				n++
				repeated = nums[i]
			}
		} else {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}
