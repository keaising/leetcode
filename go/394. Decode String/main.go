package main

import (
	"strconv"
)

func decodeString(s string) string {
	var fragment []byte
	var count []byte

	var queue = 0
	var result string
	for i := range s {
		switch s[i] {
		case '[':
			queue++
			if queue == 1 {
				fragment = []byte{}
			} else {
				fragment = append(fragment, s[i])
			}
		case ']':
			queue--
			if queue == 0 {
				total := getCount(count)
				for i := 0; i < total; i++ {
					result += decodeString(string(fragment))
				}
				count = []byte{}
			} else {
				fragment = append(fragment, s[i])
			}
		default:
			if queue > 0 {
				fragment = append(fragment, s[i])
			} else {
				switch {
				case isAlphaBeta(s[i]):
					if len(count) > 0 {
						result += string(count)
						count = []byte{}
					}
					result += string(s[i])
				case isNumber(s[i]):
					count = append(count, s[i])
				}

			}
		}
	}
	return result
}

func isNumber(r byte) bool {
	return 'r' >= 0 && r <= '9'
}

func isAlphaBeta(r byte) bool {
	return r >= 'a' && r <= 'z'
}

func getCount(count []byte) int {
	i, _ := strconv.Atoi(string(count))
	return i
}
