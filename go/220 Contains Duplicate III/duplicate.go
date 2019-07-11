package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	n := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		small := smaller(n, i+k)
		for j := i + 1; j <= small; j++ {
			if abs(nums[i], nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func smaller(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
