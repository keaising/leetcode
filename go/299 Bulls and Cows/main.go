package main

import "fmt"

func main() {
	fmt.Println(getHint("233", "233"))
}

func getHint(secret string, guess string) string {
	secretMap := make(map[string]int)
	guessMap := make(map[string]int)
	total := len(secret)
	if len(guess) < total {
		total = len(guess)
	}
	A, B := 0, 0
	for i := 0; i < total; i++ {
		if secret[i] == guess[i] {
			A += 1
		} else {
			secretMap[string(secret[i])] += 1
			guessMap[string(guess[i])] += 1
		}
	}
	if len(secret) > total {
		for j := total; j < len(secret); j++ {
			secretMap[string(secret[j])] += 1
		}
	}
	if len(guess) > total {
		for j := total; j < len(guess); j++ {
			guessMap[string(guess[j])] += 1
		}
	}
	for i := 0; i < 10; i++ {
		B += min(secretMap[fmt.Sprintf("%v", i)], guessMap[fmt.Sprintf("%v", i)])
	}
	return fmt.Sprintf("%dA%dB", A, B)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
