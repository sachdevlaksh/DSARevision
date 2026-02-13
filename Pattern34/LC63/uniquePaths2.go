/*
LeetCode Problem #63: Unique Paths II
Difficulty: Medium

You are given an m x n integer array grid where grid[i][j] = 1 represents an obstacle and grid[i][j] = 0 represents an empty cell. Return the number of unique paths from top-left to bottom-right.
*/

package LC63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= m || j >= n{
			return 0
		}

		if obstacleGrid[i][j] == 1{
			return 0
		}

		if i == m-1 && j == n-1 {
			dp[i][j] = 1
			return 1
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		dp[i][j] = dfs(i+1, j) + dfs(i, j+1)
		return dp[i][j]
	}
	return dfs(0, 0)
}
