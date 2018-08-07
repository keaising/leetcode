package foursum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	//i1,i2,i3,i4:=0,0,0,0
	ret := make([][]int, 0)
	for i1 := range nums {
		for i2 := range nums {
			for i3 := range nums {
				for i4 := range nums {
					if i4 > i3 && i3 > i2 && i2 > i1 {
						if nums[i1]+nums[i2]+nums[i3]+nums[i4] == target {
							tmp := sort1([]int{nums[i1], nums[i2], nums[i3], nums[i4]})
							if !exists(ret, tmp) {
								ret = append(ret, tmp)
							}
						}
					}
				}
			}
		}
	}
	return ret
}

func sort1(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func exists(arr [][]int, a []int) bool {
	for _, v1 := range arr {
		for k2, v2 := range v1 {
			if v2 != a[k2] {
				break
			}
			if k2 == 3 {
				return true
			}
		}
	}
	return false
}
