package main

func sortColors2(nums []int) {
	j, k := 0, len(nums)-1
	for i := 0; i < k;i++ {
		if nums[i] == 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
		if nums[i] == 2 {
			nums[i], nums[k] = nums[k], nums[i]
			k--
			i--
		}
	}
}
