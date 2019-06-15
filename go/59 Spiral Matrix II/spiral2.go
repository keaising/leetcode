package main

func generateMatrix2(n int) [][]int {
	matrix := create2DArray(n)
	if n == 0 {
		return matrix
	}

	leftBoundary := 0
	rightBoundary := n - 1
	upperBoundary := 0
	bottonBoundary := n - 1

	i := 0
	j := 0
	num := 1
	for {
		more := true

		// 0 - move right, 1 move down, 2 move left, 3 move up
		for direction := 0; direction < 4; direction++ {
			switch direction {
			case 0:
				for ; j <= rightBoundary; j++ {
					matrix[i][j] = num
					num++
				}
				j--
				i++
				upperBoundary++
			case 1:
				for ; i <= bottonBoundary; i++ {
					matrix[i][j] = num
					num++
				}
				i--
				j--
				rightBoundary--
			case 2:
				for ; j >= leftBoundary; j-- {
					matrix[i][j] = num
					num++
				}
				j++
				i--
				bottonBoundary--
			case 3:
				for ; i >= upperBoundary; i-- {
					matrix[i][j] = num
					num++
				}
				i++
				j++
				leftBoundary++
			}

			if leftBoundary > rightBoundary || upperBoundary > bottonBoundary {
				more = false
				break
			}
		}

		if !more {
			break
		}
	}

	return matrix
}

func create2DArray(n int) [][]int {
	var result [][]int
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	return result
}
