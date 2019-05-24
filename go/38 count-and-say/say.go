/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

package main

import "fmt"

func countAndSay(n int) string {
	if n > 1 {
		return say(countAndSay(n - 1))
	}
	return "1"
}

func say(str string) string {
	last := '0'
	count := 0
	ret := ""
	for i, r := range str {
		if r == last {
			count++
		} else {
			if count > 0 {
				ret = fmt.Sprintf("%s%d%s", ret, count, string(last))
				count = 1
			} else {
				count++
			}
			last = r
		}
		if i == len(str)-1 {
			if r == last {
				ret = fmt.Sprintf("%s%d%s", ret, count, string(r))
			} else {
				ret = fmt.Sprintf("%s%d%s1%s", ret, count, string(last), string(r))
			}
		}
	}
	return ret
}
