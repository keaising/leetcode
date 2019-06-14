package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}
	m := new(Matrix)
	*m = intervals
	sort.Sort(*m)

	ret := [][]int{}
	carry := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if carry[1] >= (*m)[i][0] {
			carry[1] = bigger((*m)[i][1], carry[1])
		} else {
			ret = append(ret, append([]int{}, carry...))
			carry[0] = (*m)[i][0]
			carry[1] = (*m)[i][1]
		}
	}
	ret = append(ret, append([]int{}, carry...))

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

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
