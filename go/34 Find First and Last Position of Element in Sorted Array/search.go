package main

func searchRange(nums []int, target int) []int {
	null := []int{-1, -1}
	if len(nums) == 0 {
		return null
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return []int{0, 0}
		} else {
			return null
		}
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return null
	}

	start := 0
	end := len(nums) - 1
	pStart := -1
	pEnd := -1

	for start < end {
		mid := (start + end) / 2
		if target < nums[mid] {
			end = mid - 1
		}
		if target > nums[mid] {
			start = mid + 1
		}
		if target == nums[mid] {
			left := leftSearch(target, start, mid, nums)
			right := rightSearch(target, mid, end, nums)
			return []int{left, right}
		}
	}
	return []int{pStart, pEnd}
}

func leftSearch(target, left, right int, nums []int) int {
	if left == right {
		return left
	}
	for left < right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid - 1
		}
		if target > nums[mid] {
			left = mid + 1
		}
	}
	return right + 1
}

func rightSearch(target, left, right int, nums []int) int {
	if left == right {
		return right
	}
	for left < right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		}
		if target >= nums[mid] {
			left = mid + 1
		}
	}
	return left - 1
}
