/*
LeetCode Problem #62: Unique Paths
Difficulty: Medium

There is a robot on an m x n grid. The robot is initially at the top-left corner (grid[0][0]). The robot tries to move to the bottom-right corner (grid[m-1][n-1]).
*/

package LC62

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}

		if i == m-1 && j == n-1 {
			dp[i][j] = 1
		}
		
		if dp[i][j] != -1 {
			return dp[i][j]
		}


		dp[i][j] = dfs(i+1, j) + dfs(i, j+1)
		return dp[i][j]
	}
	return dfs(0, 0)

}
