package anagram

func isAnagram(s string, t string) bool {
	dict1 := [26]byte{}
	dict2 := [26]byte{}
	for _, sc := range s {
		dict1[sc-'a']++
	}
	for _, tc := range t {
		dict2[tc-'a']++
	}
	if dict1 == dict2 {
		return true
	}
	return false
}