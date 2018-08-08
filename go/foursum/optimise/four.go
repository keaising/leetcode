package optimise

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	if len(nums) == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			return [][]int{nums}
		}
		return nil
	}
	sort.Ints(nums)
	var ret = make([][]int, 0)
	for k := range nums[0 : len(nums)-3] {
		if k == 0 || nums[k] != nums[k-1] {
			for j := k + 1; j < len(nums)-2; j++ {
				if j == k+1 || nums[j] != nums[j-1] {
					l, r := j+1, len(nums)-1
					for l < r {
						sum := nums[k] + nums[j] + nums[l] + nums[r]
						if sum == target {
							ret = append(ret, []int{nums[k], nums[j], nums[l], nums[r]})
							l++
							r--
							for l < r && nums[l] == nums[l-1] {
								l++
							}
							for l < r && nums[r] == nums[r+1] {
								r--
							}
						} else if sum < target {
							l++
						} else {
							r--
						}
					}
				}
			}
		}
	}
	return ret
}
