package main

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]int)

	for _, v := range wordDict {
		dict[v] = 0
	}

	arr := make([]bool, len(s)+1)
	arr[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dict[s[j:i]]; ok && arr[j] {
				arr[i] = true
				break
			}
		}
	}
	return arr[len(s)]
}

