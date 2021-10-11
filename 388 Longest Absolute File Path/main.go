package main

import (
	"strings"
)

func lengthLongestPath(input string) int {
	var (
		max       = 0
		pwd       []string
		pwdLength = 0
	)
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		isFile, depth, name := helper(line)
		if depth <= len(pwd)-1 {
			pwd = pwd[0:depth]
			pwdLength = len(strings.Join(pwd, "\\"))
		}
		if isFile {
			if max < len(name)+pwdLength+1 {
				if depth == 0 {
					max = len(name)
				} else {
					max = len(name) + pwdLength + 1
				}
			}
		} else {
			switch {
			case depth == len(pwd)-1:
				pwd[len(pwd)-1] = name
			case depth > len(pwd)-1:
				pwd = append(pwd, name)
			case depth < len(pwd)-1:
				pwd = pwd[0:depth]
				pwd = append(pwd, name)
			}
			pwdLength = len(strings.Join(pwd, "\\"))
		}
	}
	return max
}

func helper(str string) (bool, int, string) {
	var depth = 0
	for strings.HasPrefix(str, "\t") {
		depth += 1
		str = str[1:]
	}
	return strings.Contains(str, "."), depth, str
}
