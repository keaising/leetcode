package main

func coinChange(coins []int, amount int) int {
	var arr []int
	for i := 0; i <= amount; i++ {
		arr = append(arr, amount+1)
	}
	arr[0] = 0
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				arr[i] = min(arr[i], arr[i-coin]+1)
			}
		}
	}
	if arr[amount] == amount+1 {
		return -1
	}
	return arr[amount]
}
