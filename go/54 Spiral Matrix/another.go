package main

func spiralOrder2(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return matrix[0]
	}
	n := len(matrix[0])
	min := m
	if n < min {
		min = n
	}
	if min%2 == 0 {
		min = min / 2
	} else {
		min = min/2 + 1
	}
	res := make([]int, 0, m*n)
	for i := 0; i < min; i++ {
		helper(matrix, i, m, n, &res)
	}
	return res
}

func helper(matrix [][]int, start int, m, n int, res *[]int) {
	left := start
	right := n - start - 1
	top := start
	bottom := m - start - 1

	for i := left; i <= right; i++ {
		*res = append(*res, matrix[top][i])
	}
	for i := top + 1; i < bottom; i++ {
		*res = append(*res, matrix[i][right])
	}
	if bottom > top {
		for i := right; i >= left; i-- {
			*res = append(*res, matrix[bottom][i])
		}
	}
	if right > left {
		for i := bottom - 1; i >= top+1; i-- {
			*res = append(*res, matrix[i][left])
		}
	}
}
