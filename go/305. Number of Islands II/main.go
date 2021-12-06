package main

import "log"

func main() {
	// res := numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}})
	res := numIslands2(1, 1, [][]int{{0, 0}})
	log.Println(res)
}

func numIslands2(m int, n int, positions [][]int) []int {
	var dp [][]int
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		dp = append(dp, row)
	}

	var res []int
	var curr int
	for i, p := range positions {
		curr += 1
		order := i + 1
		if dp[p[0]][p[1]] != 0 {
			curr -= 1
			res = append(res, curr)
			continue
		}
		dp[p[0]][p[1]] = order

		for _, dir := range [][]int{
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		} {
			r := dir[0] + p[0]
			c := dir[1] + p[1]
			if r >= 0 && r < m && c >= 0 && c < n {
				if dp[r][c] != 0 && dp[r][c] != order {
					curr -= 1
					change(r, c, order, dp, m, n)
				}
			}
		}
		// log.Println(dp)
		res = append(res, curr)
	}
	return res
}

func change(pr, pc int, v int, dp [][]int, m, n int) {
	dp[pr][pc] = v
	for _, dir := range [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	} {
		r := dir[0] + pr
		c := dir[1] + pc
		if r >= 0 && r < m && c >= 0 && c < n {
			if dp[r][c] != 0 && dp[r][c] != v {
				// change all value recursively
				change(r, c, v, dp, m, n)
			}
		}
	}
}
