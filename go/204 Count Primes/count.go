package main

import "math"

func countPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++ {
		if helper(n) {
			ans++
		}
	}
	return ans
}

func helper(n int) bool {
	for i := 2; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

