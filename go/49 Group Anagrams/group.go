package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	dict := map[[26]byte]int{}
	result := [][]string{}
	for _, str := range strs {
		cache := [26]byte{}
		for _, s := range str {
			cache[s-'a']++
		}
		if idx, ok := dict[cache]; ok {
			result[idx] = append(result[idx], str)
		} else {
			result = append(result, []string{str})
			dict[cache] = len(result) - 1
		}
	}
	return result
}

func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat", "", ""})
	fmt.Println(result)
}
