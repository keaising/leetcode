package main

import "strconv"

func numDecodings(s string) int {
	if len(s) == 1 {
		if s == "0" {
			return 0
		}
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	arr := make([]int, len(s)+1)
	arr[0] = 1
	if arr[1] == '0' {
		arr[1] = 0
	} else {
		arr[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		if k, _ := strconv.Atoi(s[i-1 : i]); k >= 1 && k <= 9 {
			arr[i] += arr[i-1]
		}
		if k, _ := strconv.Atoi(s[i-2 : i]); k < 27 && k > 9 {
			arr[i] += arr[i-2]
		}
	}
	return arr[len(s)]
}
