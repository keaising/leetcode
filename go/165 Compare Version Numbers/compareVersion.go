package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	list1 := strings.Split(version1, ".")
	list2 := strings.Split(version2, ".")
	i := 0
	for ; i < len(list1) && i < len(list2); i++ {
		n1 := str2int(list1[i])
		n2 := str2int(list2[i])
		if n1 > n2 {
			return 1
		}
		if n1 < n2 {
			return -1
		}
	}
	for ; i < len(list1); i++ {
		if str2int(list1[i]) > 0 {
			return 1
		}
	}
	for ; i < len(list2); i++ {
		if str2int(list2[i]) > 0 {
			return -1
		}
	}
	return 0
}

func str2int(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}
