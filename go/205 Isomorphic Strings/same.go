package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := map[uint8]int{}
	m2 := map[uint8]int{}
	for  i := range s {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i+1
		m2[t[i]] = i+1
	}
	return true
}

