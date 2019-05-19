package main

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if len(nums) == 0 {
		return -1
	}

	for {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if start >= end {
			return -1
		}

		if nums[mid] >= nums[start] {
			//left sorted
			if nums[mid] >= target && nums[start] <= target {
				end = mid - 1
			} else if target > nums[end] && target < nums[start] {
				return -1
			} else {
				start = mid + 1
			}
		} else {
			//right sorted
			if nums[mid] <= target && nums[end] >= target {
				start = mid + 1
			} else if nums[end] < target && nums[start] > target {
				return -1
			} else {
				end = mid - 1
			}
		}
	}
}
