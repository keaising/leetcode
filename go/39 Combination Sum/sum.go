package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	ret := [][]int{}
	temp := []int{}
	sort.Ints(candidates)
	dfs(&ret, temp, candidates, target, 0)
	return ret
}

func dfs(ret *[][]int, temp []int, candidates []int, target int, sum int) {
	for i, ca := range candidates {
		if sum+ca == target {
			resTemp := make([]int, len(temp)+1)
			copy(resTemp, append(temp, ca))
			*ret = append(*ret, resTemp)
			break
		} else if sum+ca < target {
			dfs(ret, append(temp, ca), candidates[i:], target, sum+ca)
		} else {
			break
		}
	}
}
