package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	if len(nums) < 2 {
		return false
	}
	for i, n := range nums {
		if _, ok := m[n]; ok && i-m[n] <= k {
			return true
		} else {
			m[n] = i
		}
	}
	return false
}
