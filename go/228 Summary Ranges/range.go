package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	s, e := 0,0
	ans := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 == nums[i] {
			e = i
		} else {
			if s == e {
				ans = append(ans, strconv.Itoa(nums[s]))
			} else {
				ans = append(ans, fmt.Sprint(nums[s], "->", nums[e]))
			}
			s = i
			e = i
		}
	}
	if s == e {
		ans = append(ans, strconv.Itoa(nums[s]))
	} else {
		ans = append(ans, fmt.Sprint(nums[s], "->", nums[e]))
	}
	return ans
}