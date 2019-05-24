package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)
		col := make(map[byte]bool)
		sub := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			curr := board[i][j]
			if curr != '.' {
				if row[curr] {
					return false
				} else {
					row[curr] = true
				}
			}

			curr = board[j][i]
			if curr != '.' {
				if col[curr] {
					return false
				} else {
					col[curr] = true
				}
			}

			rowIndex := i/3*3 + j/3
			colIndex := i%3*3 + j%3
			curr = board[rowIndex][colIndex]
			if curr != '.' {
				if sub[curr] {
					return false
				} else {
					sub[curr] = true
				}
			}

		}
	}
	return true
}
