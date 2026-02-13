/*
LeetCode Problem #213: House Robber II
Difficulty: Medium

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
*/

package LC213

func rob(nums []int) int {
    n := len(nums)
    dp1 := make([]int, n)
    for i := range dp1 {
        dp1[i] = -1
    }
    dp2 := make([]int, n)
    for i := range dp2 {
        dp2[i] = -1
    }
    var solve func(index int, myNums []int, dp []int) int
    solve = func(index int, myNums []int, dp []int) int {
        if index >= n-1{
            return 0
        }

        if dp[index] != -1 {
            return dp[index]
        }

        //take
        take := myNums[index] + solve(index+2, myNums, dp)

        //skip
        skip := solve(index+1, myNums, dp)

        dp[index] = max(take, skip)
        return dp[index]
    }

    myNums1 := nums[:n-1]
    myNums2 := nums[1:]
	return max(solve(0, myNums1, dp1), solve(0, myNums2, dp2))
}
