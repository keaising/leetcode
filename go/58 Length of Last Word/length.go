package length

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	cur, result := 0, 0
	for _, ch := range s {
		if ch == ' ' {
			cur = 0
		} else {
			cur ++
			result = cur
		}
	}
	return result
}