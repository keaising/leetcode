package main

import "strings"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		for !isAlphanumeric(rune(s[i])) && i < j {
			i++
		}
		for !isAlphanumeric(rune(s[j])) && i < j {
			j--
		}
		if i >= j {
			return true
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
