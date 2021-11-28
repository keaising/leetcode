package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	result := visiblePoints([][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{1, 2},
		{2, 1},
	}, 0, []int{1, 1})
	log.Println(result)
}

func visiblePoints(points [][]int, angle int, location []int) int {
	var degress []float64
	var same int
	for _, p := range points {
		if p[1] == location[1] && p[0] == location[0] {
			same++
			continue
		}
		degress = append(degress, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0]))*180/math.Pi)
	}
	sort.Float64s(degress)

	for _, d := range degress {
		degress = append(degress, d+360)
	}

	var max int
	var l int
	for r := range degress {
		for (degress[r] - degress[l]) > float64(angle) {
			l += 1
			// if r > len(degress) {
			// degress[r%len(degress)] = degress[r%len(degress)] + 2*math.Phi
			// }
		}
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max + same
}
