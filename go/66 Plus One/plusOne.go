package main

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + carry
		carry = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
