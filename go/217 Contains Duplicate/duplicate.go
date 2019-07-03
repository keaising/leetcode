package main

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	if len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[i]; ok {
			return true
		} else {
			m[i] = 1
		}
	}
	return false
}
