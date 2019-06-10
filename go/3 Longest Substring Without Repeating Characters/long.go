package long

func lengthOfLongestSubstring(s string) int {
	larger := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	max, start := 0, 0
	table := [128]int{}
	for i := range table {
		table[i] = -1
	}
	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		max = larger(max, i-start+1)
	}
	return max
}
