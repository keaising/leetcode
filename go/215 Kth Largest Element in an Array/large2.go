package main

func findKthLargest(nums []int, k int) int {
	pivot := nums[0]
	n := len(nums)
	left, right := 0, len(nums)
	for i := 1; i < right; {
		if nums[i] > pivot {
			right--
			nums[i], nums[right] = nums[right], nums[i]
		} else {
			left++
			i++
		}
	}
	if n-left == k {
		return pivot
	} else if n-left > k {
		return findKthLargest(nums[right:], k)
	} else {
		return findKthLargest(nums[1:right], k-n+right-1)
	}
}
