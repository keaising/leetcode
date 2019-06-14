package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	maxPosition := 0
	for i := 1; i < len(nums); i++ {
		maxPosition = bigger(maxPosition, i-1+nums[i-1])
		if maxPosition < i {
			return false
		}
	}
	return true
}

func bigger(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return a
}
