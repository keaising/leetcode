package main

import "log"

func main() {
	res := numMatchingSubseq("abcde", []string{
		"a",
		"bb",
		"acd",
		"ace",
	})
	log.Println(res)
}

func numMatchingSubseq(s string, words []string) int {
	m := make(map[rune][]string)

	for _, word := range words {
		m[[]rune(word)[0]] = append(m[[]rune(word)[0]], word)
	}

	var res int

	for _, r := range s {

		var rs []string
		rs = append(rs, m[r]...)

		m[r] = []string{}

		for _, word := range rs {
			if len(word) == 1 {
				res += 1
				continue
			}
			m[[]rune(word)[1]] = append(m[[]rune(word)[1]], word[1:])
		}
	}
	return res
}
