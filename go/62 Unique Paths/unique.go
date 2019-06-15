package main

func uniquePaths(m int, n int) int {f
	if m == 0 || n == 0 {
		return 0
	}
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	grid[0][0] = 1
	for i := 1; i < m; i++ {
		grid[i][0] = 1
	}
	for i := 1; i < n; i++ {
		grid[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}
