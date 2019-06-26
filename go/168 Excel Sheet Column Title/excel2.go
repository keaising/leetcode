package main

func convertToTitle2(n int) string {
	ret := ""
	for n > 0 {
		r := n%26
		ret = toChar(r) + ret
		if r == 0 {
			r = 26
		}
		n = (n-r)/26
	}
	return ret
}

func toChar(n int) string {
	if n == 0 {
		n = 26
	}
	return string('A' + n - 1)
}
