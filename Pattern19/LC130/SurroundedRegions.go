/*
LeetCode Problem #130: Surrounded Regions
Difficulty: Medium

Given an m x n matrix board containing 'X' and 'O' (the letter O), capture all regions that are 4-directionally surrounded by 'X'. A region is captured by flipping all 'O's into 'X's in that surrounded region.
*/

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
