package main

import (
	"log"
)

func main() {
	log.Println(minTransfers([][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 4}, {3, 4, 5}}))
}

func minTransfers(transactions [][]int) int {
	var m = make(map[int]int) // account -> balance
	for _, t := range transactions {
		var f, t, a = t[0], t[1], t[2] // from, to, amount
		m[f] -= a
		m[t] += a
	}

	var amounts = []int{}
	for _, b := range m {
		if b != 0 { // skip balanced accounts
			amounts = append(amounts, b)
		}
	}

	return transfer(0, amounts)
}

func transfer(start int, debt []int) int {
	for start < len(debt) && debt[start] == 0 {
		start++
	}
	if start == len(debt) {
		return 0
	}

	var res = len(debt)
	for i := start + 1; i < len(debt); i++ {
		if debt[start]*debt[i] >= 0 {
			continue
		}
		debt[i] += debt[start]
		t := transfer(start+1, debt)
		if res > (t + 1) {
			res = t + 1
		}
		debt[i] -= debt[start]
		if debt[i]-debt[start] == 0 {
			break
		}
	}
	return res
}
