package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}
	closedIsland(matrix)
	fmt.Println(matrix)
}

func closedIsland(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for _, j := range []int{0, len(grid[0]) - 1} {
			if grid[i][j] == 0 {
				deleteAdjacentZeroes(grid, i, j)
			}
		}
	}
	for j := 0; j < n; j++ {
		for _, i := range []int{0, len(grid) - 1} {
			if grid[i][j] == 0 {
				deleteAdjacentZeroes(grid, i, j)
			}
		}
	}

	var res int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				res += 1
				deleteAdjacentZeroes(grid, i, j)
			}
		}
	}
	return res
}

func deleteAdjacentZeroes(grid [][]int, i, j int) {
	if !inBound(grid, i, j) {
		return
	}
	if grid[i][j] == 1 {
		return
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	q := [][]int{{i, j}}
	for len(q) > 0 {
		p := q[0]
		grid[p[0]][p[1]] = 1
		for _, d := range directions {
			r := p[0] + d[0]
			c := p[1] + d[1]
			if grid[r][c] == 0 {
				q = append(q, []int{r, c})
			}
		}
		q = q[1:]
	}
}

func inBound(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
