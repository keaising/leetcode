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

var max int
var visit [][]int

var dirs = map[int][]int{
	1: {0, 1},
	2: {0, -1},
	3: {1, 0},
	4: {-1, 0},
}

func minCost(grid [][]int) int {
	max = len(grid) * len(grid[0])
	visit = [][]int{}
	for range grid {
		var row []int
		for range grid[0] {
			row = append(row, 0)
		}
		visit = append(visit, row)
	}
	min(grid, grid[0][0], 0, 0, 0)
	return max
}

func min(grid [][]int, dir, x, y int, count int) {
	if x >= len(grid) || x < 0 {
		return
	}
	if y >= len(grid[0]) || y < 0 {
		return
	}
	if x == len(grid)-1 && y == len(grid[0])-1 {
		if max > count {
			max = count
		}
		return
	}
	if visit[x][y] == 1 {
		return
	}
	visit[x][y] = 1

	for c, d := range dirs {
		if c == grid[x][y] {
			min(grid, c, x+d[0], y+d[1], count)
		} else {
			min(grid, c, x+d[0], y+d[1], count+1)
		}
	}
	visit[x][y] = 0
}
