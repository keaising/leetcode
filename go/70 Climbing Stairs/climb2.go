package main

func climbStairs2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	cache := make([]int, n+1)
	return helper(n, &cache)
}

func helper(n int, cache *[]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if (*cache)[n] > 0 {
		return (*cache)[n]
	}
	(*cache)[n] = helper(n-1, cache) + helper(n-2, cache)
	return (*cache)[n]
}