package main

import "strconv"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	zeroes := ""
	for i := 0; i < len(a)-len(b); i++ {
		zeroes += "0"
	}
	b = zeroes + b
	carry := 0
	current := 0
	result := ""
	for i := len(a) - 1; i >= 0; i-- {
		if string(a[i]) == "1" {
			carry++
		}
		if string(b[i]) == "1" {
			carry++
		}
		current = carry % 2
		carry = carry / 2
		result = strconv.Itoa(current) + result
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}
