package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	curr := []int{}
	last := []int{}
	for i := 0; i <= rowIndex; i++ {
		last = curr
		curr = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				curr[j] = 1
			} else {
				curr[j] = last[j-1] + last[j]
			}
		}
	}
	return curr
}
