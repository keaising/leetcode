package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row, col, l := len(matrix)-1, 0, len(matrix[0])-1

	for row >= 0 && col <= l {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			col++
		} else {
			row--
		}
	}
	return false
}
