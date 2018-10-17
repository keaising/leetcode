package isValid

import (
	"strings"
)

func isValid(s string) bool {
	arr := make([]rune, 0)
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	if len(s) == 0 {
		return true
	} else if len(s)%2 == 1 {
		return false
	}
	for _, sItem := range s {
		if strings.ContainsAny("([{", string(sItem)) {
			arr = append(arr, sItem)
		} else if len(arr) == 0 {
			return false
		} else {
			last := arr[len(arr)-1]
			if last == m[sItem] {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}
	if len(arr) == 0 {
		return true
	}
	return false
}
