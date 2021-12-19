package main

import "log"

func main() {
	grid := [][]int{
		{3, 4, 3},
		{2, 2, 2},
		{2, 1, 1},
		{4, 3, 2},
		{2, 1, 4},
		{2, 4, 1},
		{3, 3, 3},
		{1, 4, 2},
		{2, 2, 1},
		{2, 1, 1},
		{3, 3, 1},
		{4, 1, 4},
		{2, 1, 4},
		{3, 2, 2},
		{3, 3, 1},
		{4, 4, 1},
		{1, 2, 2},
		{1, 1, 1},
		{1, 3, 4},
		{1, 2, 1},
		{2, 2, 4},
		{2, 1, 3},
		{1, 2, 1},
		{4, 3, 2},
		{3, 3, 4},
		{2, 2, 1},
		{3, 4, 3},
		{4, 2, 3},
		{4, 4, 4},
	}
	result := minCost(grid)
	log.Println(result)
}

func minCost(grid [][]int) int {
	max := len(grid) * len(grid[0])
	visit := [][]bool{}
	var dirs = map[int][]int{
		1: {0, 1},
		2: {0, -1},
		3: {1, 0},
		4: {-1, 0},
	}
	for range grid {
		var row []bool
		for range grid[0] {
			row = append(row, false)
		}
		visit = append(visit, row)
	}
	var q = [][]int{{0, 0, 0}}
	for len(q) > 0 {
		x, y, cost := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		if x == len(grid)-1 && y == len(grid[0])-1 {
			return cost
		}
		visit[x][y] = true
		for c, d := range dirs {
			dx, dy := x+d[0], y+d[1]
			if dx >= 0 && dx < len(grid) && dy >= 0 && dy < len(grid[0]) && !visit[dx][dy] {
				if c == grid[x][y] {
					q = append([][]int{{dx, dy, cost}}, q...)
				} else {
					q = append(q, []int{dx, dy, cost + 1})
				}
			}
		}
	}
	return max
}
