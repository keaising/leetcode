package structure

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

