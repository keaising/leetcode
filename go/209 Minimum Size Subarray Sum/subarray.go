package main

import "math"

func minSubArrayLen(s int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	from := 0
	for i := range nums {
		sum += nums[i]
		for sum >= s {
			ans = smaller(ans, i-from+1)
			sum -= nums[from]
			from++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func smaller(a, b int) int {
	if a > b {
		return b
	}
	return a
}
