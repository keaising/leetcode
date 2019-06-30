package main


func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i:=0;i<len(board);i++ {
		for j := 0; j < len(board[0]);j++ {
			ret := search(&board, word, 0, i,j)
			if ret {
				return true
			}
		}
	}
	return false
}

func search(board *([][]byte), word string, index int, row int , col int) bool {
	if len(word) == index {
		return true
	}
	if row < 0 || row > len(*board)-1 {
		return false
	}
	if col < 0 || col > len((*board)[0])-1 {
		return false
	}
	if (*board)[row][col] == '#' {
		return false
	}
	if word[index] != (*board)[row][col] {
		return false
	}
	ch := (*board)[row][col]
	(*board)[row][col] = '#'
	ret := search(board, word, index+1, row+1, col) ||
		search(board, word, index+1, row, col + 1) ||
		search(board, word, index+1, row-1, col)||
		search(board, word, index+1, row, col-1)
	(*board)[row][col] = ch
	return ret
}