package main

import (
	"fmt"
	"math"
	"strconv"
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
	if l1 > l2 {
		num1, num2 = num2, num1
	} else if num1[0] > num2[0] {
		num1, num2 = num2, num1
	}
	ret := float64(0)
	n1 := conv(num1)
	n2 := conv(num2)
	ii := float64(0)
	for i := float64(1); i <= n1; i++ {
		ret += n2
		ii = i
	}
	fmt.Println(ii)
	return FloatToString(ret)
}

func conv(num string) float64 {
	ret := float64(0)
	l := len(num)
	for i, n := range num {
		no, _ := strconv.Atoi(string(n))
		ret += math.Pow10(l-i-1) * float64(no)
	}
	return ret
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}
