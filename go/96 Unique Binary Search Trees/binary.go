package main

func numTrees(n int) int {
	ans := map[int]int{}
	ans[0] = 1
	ans[1] = 1
	ans[2] = 2
	ans[3] = 5
	for i := 4; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			sum += ans[j-1] * ans[i-j]
		}
		ans[i] = sum
	}
	return ans[n]
}
