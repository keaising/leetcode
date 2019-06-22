package main

func grayCode2(n int) []int {
	for n == 0 {
		return []int{0}
	}
	for n == 1 {
		return []int{0, 1}
	}
	ret := []int{0, 1}
	for i := 1; i < n; i++ {
		length := len(ret)
		for j := length - 1; j >= 0; j-- {
			ret = append(ret, ret[j]+(1<<uint(i)))
		}
	}
	return ret
}
