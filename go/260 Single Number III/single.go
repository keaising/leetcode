package main

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	xor &= -xor
	ans := []int{0, 0}
	for _, n := range nums {
		if xor&n == 0 {
			ans[0] ^= n
		} else {
			ans[1] ^= n
		}
	}
	return ans
}
