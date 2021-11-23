package main

import (
	"log"
)

func main() {
	log.Println(stoneGameIII([]int{1, 2, 3, 6}))
}

func stoneGameIII(stoneValue []int) string {
	var dp1, dp2, dp3 int

	for i := len(stoneValue) - 1; i >= 0; i-- {
		ans := -1 << 63
		take := stoneValue[i]
		if ans < take-dp1 {
			ans = take - dp1
		}
		if i+1 < len(stoneValue) {
			take += stoneValue[i+1]
			if ans < take-dp2 {
				ans = take - dp2
			}
		}
		if i+2 < len(stoneValue) {
			take += stoneValue[i+2]
			if ans < take-dp3 {
				ans = take - dp3
			}
		}
		dp1, dp2, dp3 = ans, dp1, dp2
	}
	if dp1 > 0 {
		return "Alice"
	}
	if dp1 < 0 {
		return "Bob"
	}
	return "Tie"
}
