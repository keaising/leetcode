package main

func permute(nums []int) [][]int {
	ret := [][]int{}
	return ret
}

func all(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{
			nums,
		}
	}


	for i, n := range nums {
		ret := all(append(nums[:i], nums[i+1:]...))
		for j, r := range ret {
			tmp := []int{}
			tmp = copy(tmp, append(n))
		}
	}
}


