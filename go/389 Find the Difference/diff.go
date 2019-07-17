package main

func findTheDifference(s string, t string) byte {
	var ans byte
	for _, ch := range []byte(s) {
		ans ^= ch
	}
	for _, ch := range []byte(t) {
		ans ^= ch
	}
	return ans
}
