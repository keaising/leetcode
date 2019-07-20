package main

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		temp := 0
		if i&1 == 1 {
			temp = 1
		}
		ans[i] = ans[i>>1] + temp
	}
	return ans
}
