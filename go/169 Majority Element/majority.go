package main

func majorityElement(nums []int) int {
	target := len(nums) / 2
	if len(nums)%2 == 1 {
		target += 1
	}
	m := map[int]int{}
	for _, n := range nums {
		m[n] += 1
		if m[n] >= target {
			return n
		}
	}
	return 0
}
