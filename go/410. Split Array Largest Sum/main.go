package main

import "log"

func main() {
	log.Println(splitArray([]int{7, 2, 5, 10, 2}, 2))
}

func splitArray(nums []int, m int) int {
	var left, right int
	for _, num := range nums {
		right += num
		if num > left {
			left = num
		}
	}
	for left < right {
		mid := (left + right) / 2
		if check(nums, m, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func check(nums []int, m, n int) bool {
	var cnt = 0
	var cache = 0
	for _, num := range nums {
		if cache+num > n {
			cache = num
			cnt++
		} else {
			cache += num
		}
	}
	cnt++
	return cnt <= m
}
