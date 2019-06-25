package main

func rotate(matrix [][]int)  {
	i,j := 0, len(matrix)-1
	for i < j {
		matrix[i],matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	for m := 0; m < len(matrix); m++ {
		for n := m; n < len(matrix); n++ {
			matrix[m][n], matrix[n][m] = matrix[n][m], matrix[m][n]
		}
	}
}
