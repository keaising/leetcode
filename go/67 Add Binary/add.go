package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	length := 0
	if len(a) > len(b) {
		length = len(a)
	} else {
		length = len(b)
	}
	fmt.Println(length)
	arr := make([]int, length+1)
	for i, r := range a {
		arr[length-1-(len(a)-1-i)] += int(r - '0')
	}
	fmt.Println(1, arr)

	for i, r := range b {
		arr[length-1-i] += int(r - '0')
	}
	fmt.Println(2, arr)
	for i := length - 1; i > 0; i-- {
		arr[i-1] += arr[i] % 2
		arr[i] = arr[i] / 2
		fmt.Println(i, arr[i])
	}
	fmt.Println(3, arr)
	result := ""
	for i, v := range arr {
		if i == 0 && v != 0 {
			result = result + "1"
		}
		result = result + strconv.Itoa(v)
	}
	return result
}

func main() {
	result := addBinary("11", "1")
	fmt.Println(result)
}
