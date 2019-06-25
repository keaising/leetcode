package main

func reverseWords(s string) string {
	ret := ""
	word := ""
	for _, ch := range s {
		if ch != ' ' {
			word = word + string(ch)
		} else if len(word) > 0 {
			if len(ret) == 0 {
				ret = word
			} else {
				ret = word + " " + ret
			}
			word = ""
		}
	}
	if len(word) > 0 {
		if len(ret) == 0 {
			ret = word
		} else {
			ret = word + " " + ret
		}
	}
	return ret
}
