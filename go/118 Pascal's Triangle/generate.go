package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{
			{1},
		}
	} else if numRows == 2 {
		return [][]int{
			{1},
			{1, 1},
		}
	}
	ret := [][]int{
		{1},
		{1, 1},
	}
	for i := 2; i < numRows; i++ {
		last := ret[i-1]
		curr := make([]int, i+1)
		curr[0] = 1
		curr[i] = 1
		for j := 1; j < i; j++ {
			curr[j] = last[j-1]+last[j]
		}
		ret = append(ret, curr)
	}
	return ret
}
