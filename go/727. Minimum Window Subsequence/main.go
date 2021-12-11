package main

import (
	"log"
	"strings"
)

func main() {
	res := minWindow("ffynmlzesdshlvugsigobutgaetsnjlizvqjdpccdylclqcbghhixpjihximvhapymfkjxyyxfwvsfyctmhwmfjyjidnfryiyajmtakisaxwglwpqaxaicuprrvxybzdxunypzofhpclqiybgniqzsdeqwrdsfjyfkgmejxfqjkmukvgygafwokeoeglanevavyrpduigitmrimtaslzboauwbluvlfqquocxrzrbvvplsivujojscytmeyjolvvyzwizpuhejsdzkfwgqdbwinkxqypaphktonqwwanapouqyjdbptqfowhemsnsl", "ntimcimzah")
	log.Println(res)
}

func minWindow(s1 string, s2 string) string {
	if len(s1) <= len(s2) {
		if s1 == s2 {
			return s1
		}
		return ""
	}
	if len(s2) == 0 {
		return ""
	}
	if len(s2) == 1 {
		if strings.Contains(s1, s2) {
			return s2
		}
		return ""
	}
	m, n := len(s1), len(s2)
	i, j := 0, 0
	start, end := -1, len(s1)
	max := len(s1)
	for i < m {
		if s1[i] == s2[j] {
			if j == n-1 {
				end = i + 1
				for j >= 0 {
					for s1[i] != s2[j] {
						i--
					}
					j--
					i--
				}
				i++
				if (end - i) < max {
					max = end - i
					start = i
				}
			}
			j += 1
		}
		i += 1
	}
	if start == -1 {
		return ""
	}
	return s1[start : start+max]
}
