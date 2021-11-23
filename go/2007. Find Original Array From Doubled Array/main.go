package main

import (
	"log"
	"sort"
)

func main() {
	result := findOriginalArray([]int{0, 0, 0, 0, 1, 1, 1, 3, 4, 2, 2, 2, 6, 8})
	log.Println(result)
}

func findOriginalArray(changed []int) []int {
	var m = make(map[int]int)
	for _, item := range changed {
		m[item] += 1
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	var result []int
	sort.Ints(keys)
	v, ok := m[0]
	if ok {
		if v%2 != 0 {
			return nil
		} else {
			for i := 0; i < v/2; i++ {
				result = append(result, 0)
			}
			keys = keys[1:]
		}
	}

	for _, key := range keys {
		if m[key] == 0 {
			continue
		}
		if m[key] > m[2*key] {
			return []int{}
		}
		for i := 0; i < m[key]; i++ {
			result = append(result, key)
		}
		m[2*key] -= m[key]
	}
	return result
}
