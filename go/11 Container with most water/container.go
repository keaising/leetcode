package container

import "fmt"

func maxArea(height []int) int {
	max, big := 0, 0
	for i, j := 0, len(height)-1; i < j; {
		fmt.Println("i-j:", i, j)
		bigger := bigger(height[i], height[j])
		if bigger {
			big = height[j]
		} else {
			big = height[i]
		}
		if area := big * (j - i); area > max {
			max = area
		}
		if bigger {
			j--
		} else {
			i++
		}
		fmt.Println("max:", max)
	}
	return max
}

func bigger(a int, b int) bool {
	if a > b {
		return true
	}
	return false
}
