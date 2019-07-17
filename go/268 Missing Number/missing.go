package main

func missingNumber(nums []int) int {
	n := len(nums)
	ans := n
	for i := 0; i < n; i++ {
		ans = ans ^ i ^ nums[i]
	}
	return ans
}
