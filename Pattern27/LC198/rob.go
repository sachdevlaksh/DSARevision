package LC198

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var solve func(index int) int
	solve = func(index int) int {
		if index >= n {
			return 0
		}

		if dp[index] != -1 {
			return dp[index]
		}

		//take
		take := nums[index] + solve(index+2)

		//skip
		skip := solve(index + 1)

		dp[index] = max(take, skip)
		return dp[index]

	}

	return solve(0)
}
