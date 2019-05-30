package main

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 == 0 || l2 == 0 {
		return "0"
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	storage := make([]int, l1+l2)
	for i := range num1 {
		for j := range num2 {
			storage[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	for i := len(storage)-1; i >1; i-- {
		storage[i-1] += storage[i] / 10
		storage[i] = storage[i] % 10
	}
	ret := strings.Trim(strings.Replace(fmt.Sprint(storage), " ", "", -1), "[]")
	if ret[0] == '0' {
		ret = ret[1:]
	}
	return ret
}
