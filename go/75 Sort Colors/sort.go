package main

func sortColors(nums []int) {
	num0, num1 := 0, 0
	for i := range nums {
		if nums[i] == 0 {
			num0++
		}
		if nums[i] == 1 {
			num1++
		}
	}
	for i := range nums {
		if i < num0 {
			nums[i] = 0
		} else if i < num0+num1 {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
