package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := strings.FieldsFunc("words and 987", func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	fmt.Println(arr)
}
