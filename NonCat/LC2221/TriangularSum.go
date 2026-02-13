/*
LeetCode Problem #2221: Find Triangular Sum of An Array
Difficulty: Medium

You are given a 0-indexed array nums of size n consisting of positive integers. You are also given a 0-indexed 2D array triangles where triangles[i] = [indexi_1, indexi_2, indexi_3].
*/

package LC2221

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([][]int, len(nums))
	dp[0] = make([]int, len(nums))
	copy(dp[0], nums)

	for i := 1; i < len(nums); i++ {
		dp[i] = make([]int, len(dp[i-1])-1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j+1]%10
		}
	}

	return dp[len(nums)-1][0] % 10

}

func TriangularSum(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	for size != 1{
		for i := 0; i < len(nums) -1; i++ {
			nums[i] = (nums[i]+nums[i+1])%10
		}
		size --
	}
	return nums[0]

}
