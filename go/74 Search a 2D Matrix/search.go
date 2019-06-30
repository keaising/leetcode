package main

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix)
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if matrix[0][0] > target || matrix[m-1][n-1] < target {
		return false
	}
	for i < j-1 {
		mid := (i + j) / 2
		if matrix[mid][n-1] > target {
			j = mid
		} else {
			i = mid
		}
	}
	if matrix[i][n-1] < target {
		return exist(matrix[j], target)
	} else {
		return exist(matrix[i], target)
	}
}

func exist(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return false
}
