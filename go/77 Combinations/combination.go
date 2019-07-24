package main

func combine(n int, k int) [][]int {
	r := [][]int{}
	if n < k || n < 1 {
		return r
	}
	helper(n, k, &r, []int{}, 1)
	return r
}

func helper(n, k int, r *[][]int, curr []int, start int) {
	if len(curr) == k {
		*r = append(*r, append([]int{}, curr...))
	} else {
		for i := start; i <= n; i++ {
			helper(n, k, r, append(curr, i), i+1)
		}
	}
}
