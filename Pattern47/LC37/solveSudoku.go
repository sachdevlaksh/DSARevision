package LC37

func solveSudoku(board [][]byte) {
	solve(board)
}

func isSafe(board *[][]byte, row, col, n int, val byte) bool {
	for i := 0; i < n; i++ {
		if (*board)[i][col] == val {
			return false
		}
		if (*board)[row][i] == val {
			return false
		}

		if (*board)[3*(row/3)+i/3][3*(col/3)+i%3] == val {
			return false
		}

	}

	return true

}

func solve(board [][]byte) bool {
	n := len(board)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == '.' {
				for i := 1; i <= 9; i++ {
					if isSafe(&board, row, col, n, byte('0'+i)) {
						board[row][col] = byte('0' + i)
						if solve(board) {
							return true
						} else {
							board[row][col] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
