package max

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1 << 31
	}
	if len(nums) == 1 {
		return nums[0]
	}
	cur, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v
		} else {
			cur += v
		}
		if cur > max {
			max = cur
		}
	}
	return max
}