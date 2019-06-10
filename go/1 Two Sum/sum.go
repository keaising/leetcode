package sum

func twoSum(nums []int, target int) []int {
	for i, m := range nums {
		for j, n := range nums {
			if i == j {
				continue
			}
			if m+n == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
