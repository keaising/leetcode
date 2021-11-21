package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	result := canReorderDoubled([]int{2, 2, 4, 4, -2, -4})
	log.Println(result)
	result = canReorderDoubled([]int{0, 2, 4, 0, -2, -4})
	log.Println(result)
	result = canReorderDoubled([]int{4, -2, 2, -4})
	log.Println(result)
}

func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)
	for _, i := range arr {
		m[i] += 1
	}
	var newArr []int
	for k := range m {
		newArr = append(newArr, k)
	}
	sort.SliceStable(newArr, func(i, j int) bool {
		return math.Abs(float64(newArr[i])) < math.Abs(float64(newArr[j]))
	})
	log.Println(newArr)
	for _, i := range newArr {
		if m[i] == 0 {
			continue
		}
		if m[i] > m[i*2] {
			log.Println(i, i*2, m[i], m[i*2])
			return false
		}
		m[i*2] -= m[i]
	}
	return true
}
