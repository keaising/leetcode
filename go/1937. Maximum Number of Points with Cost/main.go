package main

import "log"

func main() {
	p := [][]int{
		{1, 5},
		{2, 3},
		{4, 2},
	}
	result := maxPoints(p)
	log.Println(result)
}

func maxPoints(points [][]int) int64 {
	var max = bigger(points[0]...)
	for i := 1; i < len(points); i++ {
		dp := points[i-1]
		for j := 1; j < len(points[0]); j++ {
			dp[j] = bigger(dp[j-1]-1, dp[j])
		}
		for j := len(points[0]) - 2; j > -1; j-- {
			dp[j] = bigger(dp[j+1]-1, dp[j])
		}

		for j := 0; j < len(points[0]); j++ {
			points[i][j] += dp[j]
			if max < points[i][j] {
				max = points[i][j]
			}
		}
	}

	return int64(max)
}

func bigger(arr ...int) int {
	var max = -1 << 63
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}
