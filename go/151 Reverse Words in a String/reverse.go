package main

func reverseWords(s string) string {
	stack := []string{}
	word := ""
	for _, ch := range s {
		if ch != ' ' {
			word = word + string(ch)
		} else if len(word) > 0 {
			stack = append(stack, word)
			word = ""
		}
	}
	if len(word) > 0 {
		stack = append(stack, word)
	}
	if len(stack) == 0 {
		return ""
	}
	if len(stack) == 1 {
		return stack[0]
	} else {
		ret := stack[len(stack)-1]
		for i := len(stack) - 2; i >= 0; i-- {
			ret += " " + stack[i]
		}
		return ret
	}
}