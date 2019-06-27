/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

package rob

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return bigger(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return bigger(nums[0], nums[1], nums[0]+nums[2])
	}
	nums[2] = nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		nums[i] = bigger(nums[i]+nums[i-2], nums[i]+nums[i-3])
	}
	return bigger(nums[len(nums)-1], nums[len(nums)-2])
}

func bigger(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	for _, n := range nums {
		if n > ret {
			ret = n
		}
	}
	return ret
}
