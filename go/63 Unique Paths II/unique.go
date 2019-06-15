package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	array := make([]int, len(obstacleGrid[0]))
	flag := true
	for i := range array {
		if obstacleGrid[0][i] == 1 {
			flag = false
		}
		if flag {
			array[i] = 1
		} else {
			array[i] = 0
		}
	}
	flag = array[0] == 1
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0  {
				if !flag {
					array[j] = 0
					continue
				}
				if obstacleGrid[i][j] == 0 {
					array[j] = 1
					continue
				}
				if obstacleGrid[i][j] == 1 {
					array[j] = 0
					flag = false
					continue
				}
			}
			if obstacleGrid[i][j] == 0 {
				array[j] += array[j-1]
			} else {
				array[j] = 0
			}
		}
	}
	return array[n-1]
}
