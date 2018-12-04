package insert

func searchInsert(nums []int, target int) int {
	if len(nums) == 1{
		if nums[0] > target {
			return 0
		}
		return 1
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < target {
			return i
		}
	}
	return len(nums)
}
