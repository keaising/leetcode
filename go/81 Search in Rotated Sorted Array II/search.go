package main

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start, mid, end := 0, (len(nums)-1)/2, len(nums)-1
	for start <= end {
		if nums[mid] == target {
			return true
			} else {
				if nums[mid] < nums[end] || nums[mid] < nums[start] {
					if target == nums[end] {
						return true
					}
					if target < nums[mid] || target > nums[end] {
						end = mid - 1
					} else {
						start = mid + 1
					}
				} else if nums[mid] > nums[start] || nums[mid] > nums[end] {
					if target == nums[start] {
						return true
					}
					if target < nums[start] || target > nums[mid] {
						start = mid + 1
					} else {
						end = mid - 1
					}
			} else {
				end--
			}
			mid = (start + end) / 2
		}
	}
	return false
}
