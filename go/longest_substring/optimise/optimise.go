package optimise

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	myMap := make(map[int32]int)
	max := 0
	start := 0
	for k, c := range s {
		i, ok := myMap[c]
		if ok && i >= start {
			start = i + 1
		}
		myMap[c] = k
		if k-start+1 > max {
			max = k - start + 1
		}
	}
	return max
}
