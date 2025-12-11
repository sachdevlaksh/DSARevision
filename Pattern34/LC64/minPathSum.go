package LC64

import "math"

func minPathSum(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])


	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(rown, coln int) int

	dfs = func(rown, coln int) int {

		if  rown >= rows || coln >= cols {
			return math.MaxInt
		}
		if rown == rows-1 && coln == cols-1 {
			return 	grid[rown][coln]
		}
		if dp[rown][coln] != -1 {
			return dp[rown][coln]
        }

		right :=	dfs(rown+1, coln)
		down := dfs(rown, coln+1)
		best := down
		if right < down{
			best = right
		}
		dp[rown][coln] = best + grid[rown][coln]
		return dp[rown][coln]
	}
	return dfs(0, 0)
}
