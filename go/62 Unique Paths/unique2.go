package main

func uniquePaths2(m int, n int) int {
	if m > n {
		m, n = n, m
	}
	array := make([]int, n)
	for i := range array {
		array[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			array[j] += array[j-1]
		}
	}
	return array[n-1]
}
