/*
LeetCode Problem #221: Maximal Square
Difficulty: Medium

Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
*/

package LC221

func maximalSquare(matrix [][]byte) int {

	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	maxSide := 0

	var dfs func(r, c int) int

	dfs = func(r, c int) int {

		if r >= n || c >= m {
			return 0
		}

		if dp[r][c] != -1 {
			return dp[r][c]
		}

		if matrix[r][c] == '0' {
			dp[r][c] = 0
		} else {
			right := dfs(r, c+1)
			down := dfs(r+1, c)
			diag := dfs(r+1, c+1)
			dp[r][c] = 1 + min3(right, down, diag)
		}

		if dp[r][c] > maxSide {
			maxSide = dp[r][c]
		}

		return dp[r][c]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j)
		}
	}
	return maxSide * maxSide

}

func min3(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}

	if b < c {
		return b
	}
	return c
}
