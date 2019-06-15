package main

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if m == 1 {
		return sum(grid[0])
	}
	if n == 1 {
		return sumGrid(grid)
	}
	for i := m - 2; i >= 0; i-- {
		grid[i][n-1] += grid[i+1][n-1]
	}
	for i := n - 2; i >= 0; i-- {
		grid[m-1][i] += grid[m-1][i+1]
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			grid[i][j] += smaller(grid[i+1][j], grid[i][j+1])
		}
	}
	return grid[0][0]
}

func smaller(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func sum(array []int) int {
	result := 0
	for _, a := range array {
		result += a
	}
	return result
}

func sumGrid(grid [][]int) int {
	result := 0
	for _, array := range grid {
		result += array[0]
	}
	return result
}
