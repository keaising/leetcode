package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	m := new(Matrix)
	*m = intervals
	sort.Sort(*m)

	ret := [][]int{}
	carry := intervals[0]
	for _, m := range *m {
		if carry[1] <= m[0] {
			carry[1] = m[1]
		} else {
			ret = append(ret, append([]int{}, carry...))
			carry[0] = m[0]
			carry[1] = m[1]
		}
	}
	if carry[0] == (*m)[len(*m)-1][0] {
		ret = append(ret, append([]int{}, carry...))
	}
	return ret
}

type Matrix [][]int

func (m Matrix) Len() int {
	return len(m)
}

func (m Matrix) Less(a, b int) bool {
	for x := range m[a] {
		if m[a][x] == m[b][x] {
			continue
		}
		return m[a][x] < m[b][x]
	}
	return false
}

func (m Matrix) Swap(a, b int) {
	m[a], m[b] = m[b], m[a]
}
