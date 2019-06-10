package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{0, 0}
	}
	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{0, 0}
}
