/*
LeetCode Problem #300: Longest Increasing Subsequence
Difficulty: Medium

Given an integer array nums, return the length of the longest strictly increasing subsequence.
*/

package LC300

func lengthOfLIS(nums []int) int {

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(nums []int, curr, prev, n int, dp *[][]int) int {

	dfs = func(nums []int, curr, prev, n int, dp *[][]int) int {
		if curr == n {
			return 0
		}
		if (*dp)[curr][prev+1] != -1 {
			return (*dp)[curr][prev+1]
		}
		include := 0
		if prev == -1 || nums[curr] > nums[prev+1] {
			include = 1 + dfs(nums, curr+1, curr, n, dp)
		}
		exclude := 0 + dfs(nums, curr+1, prev, n, dp)

		(*dp)[curr][prev+1] = max(include, exclude)
		return (*dp)[curr][prev+1]
	}
	return dfs(nums, 0, -1, len(nums), &dp)
}
