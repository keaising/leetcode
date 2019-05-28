package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ret := [][]int{}
	tmp := []int{}

	dfs(&ret, tmp, target, candidates, 0)
	return ret
}

func dfs(ret *[][]int, temp []int, target int, candidates []int, sum int) {
	pre := -1
	for i, ca := range candidates {
		if pre != -1 && ca == pre {
			continue
		}
		if sum+ca == target {
			resTemp := make([]int, len(temp)+1)
			copy(resTemp, append(temp, ca))
			*ret = append(*ret, resTemp)
			break
		} else if sum+ca < target {
			dfs(ret, append(temp, ca), target, candidates[i+1:], sum+ca)
			pre = ca
		} else {
			break
		}
	}
}
