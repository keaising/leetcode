package LongestPalindromic

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range s {
		for j := range s {
			if i >= j {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}
	max := 1
	rf, rt := 0, 0
	for k := 1; k < len(s); k++ {
		for i := 1; k+i < len(s); i++ {
			j := i + k
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] {
					if k+1 > max {
						max = k + 1
						rf = i
						rt = j
					}
				}
			}
		}
	}
	return s[rf : rt+1]
}
