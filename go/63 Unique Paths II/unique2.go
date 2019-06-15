package main

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	flag := true
	for i := range obstacleGrid[0] {
		if obstacleGrid[0][i] == 1 {
			flag = false
		}
		if flag {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}
	flag = obstacleGrid[0][0] == 1
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				if !flag {
					obstacleGrid[i][j] = 0
					continue
				}
				if obstacleGrid[i][j] == 0 {
					obstacleGrid[i][j] = 1
					continue
				}
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
					flag = false
					continue
				}
			}
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
