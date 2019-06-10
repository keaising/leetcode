package LongestPalindromic

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return ""
	}
	max := 0
	res := ""
	for i := 0; i < len(s) -1; i++ {
		str := lengthOfPalindrome(s, i, i)
		if max < len(str) {
			max = len(str)
			res = str
		}
		str = lengthOfPalindrome(s, i, i+ 1)
		if max < len(str) {
			max = len(str)
			res = str
		}
	}
	return res
}
func lengthOfPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
