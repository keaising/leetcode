package main

func countPrimes2(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isNotPrime[j] = true
		}
	}
	ans := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			ans++
		}
	}
	return ans
}
