package main

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	list := []string{}
	for i := 1; i <= n; i++ {
		list = append(list, strconv.Itoa(i))
	}
	k--
	fact := []int{1}
	for i := 1; i <= n; i++ {
		fact = append(fact, fact[len(fact)-1]*i)
	}
	var ans []string
	for i := n; i >= 1; i-- {
		f := fact[i-1]
		y := k / f
		k = k % f
		ans = append(ans, list[y])
		list = remove(list, y)
	}
	return strings.Join(ans, "")
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
