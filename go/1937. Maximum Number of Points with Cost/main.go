package main

import "log"

func main() {
	p := [][]int{
		{1, 5},
		{2, 3},
		{4, 2},
	}
	result := maxPoints(p)
	log.Println(result)
}

func maxPoints(points [][]int) int64 {
	if len(points) == 0 {
		return 0
	}
	var n = len(points)
	var m = len(points[0])
	var max int
	var row []int
	var newRow []int
	for _, item := range points[0] {
		row = append(row, item)
		newRow = append(newRow, item)
		if max < item {
			max = item
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			var rowMax int
			for k := 0; k < m; k++ {
				if row[k]-abs(k, j) > rowMax {
					rowMax = row[k] - abs(k, j)
					// log.Println(i, j, "k:", k, "abs:", abs(k, j), "row_max:", rowMax)
				}
			}
			newRow[j] = rowMax + points[i][j]
		}
		for l := range row {
			row[l] = newRow[l]
		}
	}
	for _, item := range row {
		if item > max {
			max = item
		}
	}
	return int64(max)
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
