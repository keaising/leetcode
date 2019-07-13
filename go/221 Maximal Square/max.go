package main

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + int(matrix[i][j]-'0')
			}
			ans = bigger(ans, dp[i][j])
		}
	}
	return ans * ans
}

func min(a, b, c int) int {
	switch {
	case a >= b || a >= c:
		return smaller(b, c)
	case b >= c || b >= a:
		return smaller(a, c)
	case c >= a || c >= b:
		return smaller(a, b)
	default:
		return 0
	}
}

func smaller(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func bigger(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
