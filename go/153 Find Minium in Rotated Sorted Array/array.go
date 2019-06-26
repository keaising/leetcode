/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}
