package LC51

import "strings"

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}

	}
	return solve(0, board, n)

}

func solve(col int, board [][]byte, n int) [][]string {
	ans := make([][]string, 0)
	if col == n {
		return [][]string{addSolution(board, n)}
	}

	for row := 0; row < n; row++ {
		if isSafe(row, col, board, n) {
			(board)[row][col] = 'Q'
			ans = append(ans, solve(col+1, board, n)...)
			(board)[row][col] = '.'
		}
	}
	return ans
}

func addSolution(board [][]byte, n int) []string {

	ans := make([]string, 0)

	for i := 0; i < n; i++ {
		var temp strings.Builder
		for j := 0; j < n; j++ {
			temp.WriteByte((board)[i][j])
		}
		ans = append(ans, temp.String())
	}

	return ans

}
func isSafe(row int, col int, board [][]byte, n int) bool {
	x := row
	y := col

	for y >= 0 {
		if (board)[x][y] == 'Q' {
			return false
		}
		y--
	}
	x = row
	y = col
	for x >= 0 && y >= 0 {
		if (board)[x][y] == 'Q' {
			return false
		}
		y--
		x--
	}

	x = row
	y = col
	for x < n && y >= 0 {
		if (board)[x][y] == 'Q' {
			return false
		}
		y--
		x++
	}
	return true

}

