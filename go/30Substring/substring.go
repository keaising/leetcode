package substring

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	length := len(words[0])
	m1 := make(map[string]int, 0)
	m2 := make(map[string]int, 0)
	ret := make([]int, 0)
	for _, w := range words {
		if len(w) != length {
			return nil
		}
		m1[w]++
	}
	slen := len(s)
	for i := range s {
		for j := range words {
			if i+length*(j) > slen-length {
				m2 = make(map[string]int, 0)
				break
			}
			sub := s[i+length*j : i+length*(j+1)]
			if _, ok := m1[sub]; !ok {
				m2 = make(map[string]int, 0)
				break
			}
			m2[sub]++
		}
		if mapEqual(m1, m2) {
			ret = append(ret, i)
		}
		m2 = make(map[string]int, 0)
	}
	return ret
}

func mapEqual(m1, m2 map[string]int) bool {
	for k, v := range m1 {
		if m, ok := m2[k]; !ok || m != v {
			return false
		}
	}
	return true
}
