package main

func generateMatrix(n int) [][]int {
	up, down, left, right := 0, n-1, 0, n-1
	k, count := 0, n*n
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	for {
		for col := left; col <= right; col++ {
			k++
			result[up][col] = k
		}
		up++
		if k >= count {
			break
		}
		for row := up; row <= down; row++ {
			k++
			result[row][right] = k
		}
		right--
		if k >= count {
			break
		}
		for col := right; col >= left; col-- {
			k++
			result[down][col] = k
		}
		down--
		if k >= count {
			break
		}
		for row := down; row >= up; row-- {
			k++
			result[row][left] = k
		}
		left++
		if k >= count {
			break
		}
	}
	return result
}
