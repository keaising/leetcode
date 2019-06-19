package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ret := [][]int{{}}
	j, size := 0, 0
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			j = size
		} else {
			j = 0
		}
		size = len(ret)
		for j < size {
			ret = append(ret, append(append([]int{}, ret[j]...), nums[i]))
			j++
		}
	}
	return ret
}
