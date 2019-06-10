package next

func nextPermutation(nums []int) {

	l := len(nums)
	i := l - 1
	// Break on finding the first number where the previous number is lower
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}

	nextHigher := -1

	// find next higher number
	for j := i; i > 0 && j < l; j++ {
		if nums[j] > nums[i-1] && (nextHigher == -1 || nums[j] <= nums[nextHigher]) {
			nextHigher = j
		}
	}

	if nextHigher != -1 {
		nums[i-1], nums[nextHigher] = nums[nextHigher], nums[i-1]
	}

	//Reverse the numbers
	s := i
	e := l - 1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}
