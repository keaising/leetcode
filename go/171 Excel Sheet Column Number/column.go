package main

func titleToNumber(s string) int {
	ret := 0
	for _, r := range s {
		ret = ret*26 + toInt(r)
	}
	return ret
}

func toInt(r rune) int {
	return int(r-'A') + 1
}
