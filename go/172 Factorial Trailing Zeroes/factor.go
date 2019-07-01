package main

func trailingZeroes(n int) int {
	ans, temp:=0,0
	for n/5 > 0 {
		ans += n/5
		temp = n/5
		n = temp
	}
	return ans
}
