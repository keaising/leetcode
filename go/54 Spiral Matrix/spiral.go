package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return nil
	}
	var result []int
	m, n := len(matrix), len(matrix[0])
	u, d, l, r, k := 0, m-1, 0, n-1, 0
	for k < m*n {
		for col := l; col <= r; col++ {
			result = append(result, matrix[u][col])
			k++
		}
		if k >= m*n {
			break
		}
		u++
		for row := u; row <= d; row ++ {
			result = append(result, matrix[row][r])
			k++
		}
		if k >= m*n {
			break
		}
		r--
		for col := r; col >= l; col-- {
			result = append(result, matrix[d][col])
			k++
		}
		if k >= m*n {
			break
		}
		d--
		for row := d; row >= u; row -- {
			result = append(result, matrix[row][l])
			k++
		}
		l++
		if k >= m*n {
			break
		}
	}
	return result
}
