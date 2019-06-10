package threesum

import (
	"sort"
)

func threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == target {
			return [][]int{nums}
		}
		return nil
	}

	ret := make([][]int, 0)
	for k, v := range nums[0 : len(nums)-3] {
		if k == 0 || nums[k] != nums[k-1] {
			sum := target - v
			i, j := k+1, len(nums)-1
			for i < j {
				if nums[i]+nums[j] == sum {
					ret = append(ret, []int{nums[k], nums[i], nums[j]})
					i++
					j--
					for i < j && nums[i] == nums[i+1] {
						i++
					}
					for i < j && nums[j-1] == nums[j] {
						j--
					}
				} else if nums[i]+nums[j] < sum {
					i++
				} else {
					j--
				}
			}
		}
	}

	return ret
}
