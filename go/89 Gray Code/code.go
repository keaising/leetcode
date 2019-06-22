package main

func grayCode(n int) []int {
	for n == 0 {
		return []int{0}
	}
	for n == 1 {
		return []int{0, 1}
	}
	minus := grayCode(n - 1)
	ret := append([]int{}, minus...)
	for i := len(minus) - 1; i >= 0; i-- {
		ret = append(ret, minus[i]+(1<<uint(n-1)))
	}
	return ret
}
