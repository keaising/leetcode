package main


func wordBreak2(s string, wordDict []string) bool {
	// word dict as a map
	dict := make(map[string]int)

	for _, v := range wordDict {
		dict[v] = 0
	}

	// bool dynamic reference
	ref := make([]bool, len(s))

	for i := len(s)-1; i >= 0; i-- {
		for j := i; j <= len(s); j++ {

			if _, ok := dict[s[i:j]]; ok {
				//fmt.Println(s[i:j],i,j)
				if j == len(s) {
					ref[i] = true
					continue
				}
				if ref[j] {
					ref[i] = true
				}
			}
		}
	}

	//fmt.Println(ref)
	return ref[0]
}