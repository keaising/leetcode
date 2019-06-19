package main

func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	a := nums[n-k:]
	b := nums[:n-k]
	ret := append(a, b...)
	copy(nums, ret)
}

