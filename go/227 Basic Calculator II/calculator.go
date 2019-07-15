package main

func calculate(s string) int {
	ans, cur := 0, 0
	stack := []int{}
	op := '+'
	for i, ch := range s {
		if ch-'0' >= 0 && ch-'0' <= 9 {
			cur = cur*10 + int(ch-'0')
		}
		if ch == ' ' {
			continue
		}
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			if op == '+' {
				stack = append(stack, cur)
			} else if op == '-' {
				stack = append(stack, -cur)
			} else if op == '*' {
				t := cur * stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, t)
			} else if op == '/' {
				t := stack[len(stack)-1] / cur
				stack = stack[:len(stack)-1]
				stack = append(stack, t)
			}
			op = rune(ch)
			cur = 0
		}
	}

	for _, n := range stack {
		ans += n
	}
	return ans
}
