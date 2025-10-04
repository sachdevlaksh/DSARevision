package LC130

func solve(board [][]byte) {

	rows := len(board)
	cols := len(board[0])

	var dfs func(int, int)

	dfs = func(r int, c int) {
		if r >= rows || c >= cols || r < 0 || c < 0 || board[r][c] != 'O' {
			return
		}

		board[r][c] = 'T'

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)

	}
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}

	}
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(rows-1, j)
		}

	}

	for i := 0; i< rows ; i++{
		for j := 0; j < cols; j++{
			if board[i][j] == 'T'{
				board[i][j] = 'O'
			}else if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
		}
	}
}
