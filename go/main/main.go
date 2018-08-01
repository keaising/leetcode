package main

import "fmt"

func main() {
	num := 7
	switch {
	case num > 3:
		fmt.Println("num > 3")
	case num > 4:
		fmt.Println("num > 4")
	default:
		fmt.Println("default")
	}
}
