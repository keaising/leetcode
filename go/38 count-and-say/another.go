package main

import "strconv"

func countBefore(s string) string {
	resStr := ""
	repeatCount := 1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			repeatCount++
		} else {
			resStr += strconv.Itoa(repeatCount) + string(s[i])
			repeatCount = 1
		}
	}
	return resStr
}

func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}
	return countBefore(countAndSay(n - 1))
}
