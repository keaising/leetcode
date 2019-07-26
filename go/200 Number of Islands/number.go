package main

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(&grid, i, j, len(grid), len(grid[0]))
				count++
			}
		}
	}
	return count
}

func dfs(grid *[][]byte, i, j int, r, c int) {
	if i >= 0 && i < r && j >= 0 && j < c {
		if (*grid)[i][j] == '1' {
			(*grid)[i][j] = '0'
			dfs(grid, i+1, j, r, c)
			dfs(grid, i, j-1, r, c)
			dfs(grid, i, j+1, r, c)
			dfs(grid, i-1, j, r, c)
		}
	}
}
