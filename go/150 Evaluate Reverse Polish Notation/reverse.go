package main

import "strconv"

func evalRPN(tokens []string) int {
	s := []int{}
	for _, token := range tokens {
		switch {
		case token == "+": {
			cal1 := s[len(s)-1]
			cal2 := s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, cal1+cal2)
		}
		case token == "-": {
			cal1 := s[len(s)-1]
			cal2 := s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, cal2-cal1)
		}
		case token == "*": {
			cal1 := s[len(s)-1]
			cal2 := s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, cal1*cal2)
		}
		case token == "/": {
			cal1 := s[len(s)-1]
			cal2 := s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, cal2/cal1)
		}
		default:
			i, _ := strconv.Atoi(token)
			s = append(s, i)
		}
	}
	return s[0]
}
