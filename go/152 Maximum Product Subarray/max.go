/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret, max, min := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		max = bigger(nums[i], nums[i]*max)
		min = smaller(nums[i], nums[i]*min)
		ret = bigger(ret, max)
	}
	return ret
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func smaller(a, b int) int {
	if a > b {
		return b
	}
	return a
}
