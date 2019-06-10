package letterCombinations

import "strings"

func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	m := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	for _, d := range digits {
		if len(ret) == 0 {
			ret = strings.Split(m[d], "")
			continue
		}
		word := m[d]
		tmp := make([]string, 0)
		for _, r := range ret {
			for _, w := range word {
				tmp = append(tmp, r+string(w))
			}
		}
		ret = tmp
	}
	return ret
}
