package main

func singleNumber(nums []int) int {
	dict := map[int]int{}
	for _, num := range  nums {
		if dict[num] > 0 {
			delete(dict, num)
		} else {
			dict[num]++
		}
	}
	for k := range dict{
		return k
	}
	return 0
}