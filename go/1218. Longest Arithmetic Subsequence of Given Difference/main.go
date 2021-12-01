package main

import "log"

func main() {
	res := longestSubsequence([]int{0, 0, 0, 0}, 0)
	log.Println(res)
}

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	max := 1

	for _, n := range arr {
		curr := dp[n-difference]
		curr++
		if curr > max {
			max = curr
		}
		dp[n] = curr
	}
	return max
}
