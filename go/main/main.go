package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 12
	if v,ok:=m["b"];ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key not found")
	}
}
