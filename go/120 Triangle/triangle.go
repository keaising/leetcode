package main

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = triangle[i][j] + smaller(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func smaller(a, b int) int {
	if a > b {
		return b
	}
	return a
}
