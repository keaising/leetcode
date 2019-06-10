package remove

func removeElement(nums []int, val int) int {
	index := 0
	for i := range nums {
		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != val {
					nums[i] = nums[j]
					nums[j] = val
					index++
					break
				}
			}
		} else {
			index++
		}
	}
	nums = nums[0:index]
	return index
}
