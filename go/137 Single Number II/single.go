package main

func singleNumber(nums []int) int {
	dict := map[int]int{}
	for _, n := range nums {
		dict[n]++
	}
	for k, v := range dict {
		if v == 1 {
			return k
		}
	}
	return 0
}
