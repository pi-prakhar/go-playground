package main

func exist(board [][]byte, word string) bool {
	for col := 0; col < len(board); col++ {
		for row := 0; row < len(board[col]); row++ {
			if board[col][row] == word[0] {
				if backtrack(board, col, row, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func backtrack(board [][]byte, col int, row int, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if col < 0 || col >= len(board) || row < 0 || row >= len(board[col]) || board[col][row] != word[index] {
		return false
	}

	temp := board[col][row]
	board[col][row] = '#'

	found := backtrack(board, col-1, row, word, index+1) || // Up
		backtrack(board, col+1, row, word, index+1) || // Down
		backtrack(board, col, row-1, word, index+1) || // Left
		backtrack(board, col, row+1, word, index+1) // Right

	board[col][row] = temp

	return found
}
