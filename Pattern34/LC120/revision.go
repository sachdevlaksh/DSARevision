/*
LeetCode Problem #120: Triangle
Difficulty: Medium

Given a triangle array, return the minimum path sum from top to bottom. For each step, you may move to an adjacent number of the row below. More formally, if you are on index j in the current row.
*/

package LC120

import "math"

func minimumTotalRev(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == len(triangle) -1 {
			return triangle[i][j]
		}

		if dp[i][j] != math.MinInt {
			return dp[i][j]
		}

		down := dfs(i+1, j)
		downRight := dfs(i+1, j+1)

	dp[i][j] = 	triangle[i][j] + min(down, downRight)
		return dp[i][j]
	}

	return dfs(0, 0)
}
