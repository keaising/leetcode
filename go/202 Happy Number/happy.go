package main

func isHappy(n int) bool {
	dict := map[int]int{}
	for {
		n = happy(n)
		if n == 1 {
			return true
		}
		if _, ok := dict[n]; ok {
			return false
		}
		dict[n]++
	}
}

func happy(i int) int {
	ans := 0
	for i > 0 {
		ans += (i % 10) * (i % 10)
		i = i / 10
	}
	return ans
}
