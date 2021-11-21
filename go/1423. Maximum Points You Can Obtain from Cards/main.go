package main

import "log"

func main() {
	result := maxScore([]int{1, 2, 3, 4, 5}, 3)
	log.Println(result)
}

func maxScore(cardPoints []int, k int) int {
	var sum = 0
	if k >= len(cardPoints) {
		for _, c := range cardPoints {
			sum += c
		}
		return sum
	}
	var (
		left  = make([]int, k+1)
		right = make([]int, k+1)
	)
	sum = 0
	left[0] = 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
		left[i+1] = sum
	}
	sum = 0
	right[0] = 0
	for i := 0; i < k; i++ {
		sum += cardPoints[len(cardPoints)-1-i]
		right[i+1] = sum
	}
	log.Println(left, right)
	sum = 0
	for i := 0; i <= k; i++ {
		if sum < left[i]+right[k-i] {
			sum = left[i] + right[k-i]
		}
	}
	return sum
}
