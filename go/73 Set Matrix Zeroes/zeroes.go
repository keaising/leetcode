package main

func setZeroes(matrix [][]int) {
	r, c := false, false
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
				if i == 0 {
					r = true
				}
				if j == 0 {
					c = true
				}
			}
		}
	}

	for i := len(matrix) - 1; i > 0; i-- {
		for j := len(matrix[0]) - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if c {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
	if r {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
}
