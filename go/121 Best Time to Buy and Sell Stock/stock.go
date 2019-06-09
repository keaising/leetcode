package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	max := 0
	in := prices[0]
	for _, price := range prices {
		if price >= in {
			if price - in  > max {
				max = price - in
			}
		} else {
			in = price
		}
	}
	return max
}