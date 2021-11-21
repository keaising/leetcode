package main

import "log"

func main() {
	var matrix [][]int
	matrix = append(matrix, []int{0, 1, 1, 1})
	matrix = append(matrix, []int{1, 1, 1, 1})
	matrix = append(matrix, []int{0, 1, 1, 1})

	result := countSquares(matrix)
	log.Println(result)
}

func countSquares(matrix [][]int) int {
	if matrix == nil {
		return 0
	}
	result := 0
	row := len(matrix)
	col := len(matrix[0])
	dp := [][]int{}
	for i := 0; i < row; i++ {
		var r []int
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {

				r = append(r, 0)
			} else {

				r = append(r, 1)
			}
		}
		dp = append(dp, r)
	}
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			down := 0
			right := 0
			downRight := 0
			if i+1 < row {
				down = dp[i+1][j]
			}
			if j+1 < col {
				right = dp[i][j+1]
			}
			if i+1 < row && j+1 < col {
				downRight = dp[i+1][j+1]
			}
			// choose down as flag
			if down > right {
				down = right
			}
			if down > downRight {
				down = downRight
			}
			down += 1
			result += down
			dp[i][j] = down
		}
	}
	return result
}
