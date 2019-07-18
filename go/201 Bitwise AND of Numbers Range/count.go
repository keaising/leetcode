package main

func rangeBitwiseAnd(m int, n int) int {
	count := 0
	for n != m {
		n >>= 1
		m >>= 1
		count++
	}
	return m << uint(count)
}
