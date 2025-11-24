package LC416

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	target := sum / 2
	var dfs func(index int, sum int) bool
	dfs = func(index int, sum int) bool {

		if index >= n {
			return false
		}
		if sum-target > 0 {
			return false
		}
		if sum-target == 0 {
			return true
		}
		return dfs(index+1, sum+nums[index]) || dfs(index+1, sum)
	}
	if rem := sum % 2; rem == 1 {
		return false
	} else {
		return dfs(0, 0)
	}

}
func canPartitionMem(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if rem := sum % 2; rem == 1 {
		return false
	}
	target := sum / 2

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i, v := range dp {
		for j := range v {
			dp[i][j] = -1
		}
	}

	var dfs func(index int, curr int) bool
	dfs = func(index int, curr int) bool {

		if curr-target == 0 {
			return true
		}

		if index >= n || curr-target > 0 {
			return false
		}

		if dp[index][curr] != -1 {
			return dp[index][target] == 1
		}
		if dfs(index+1, curr+nums[index]) || dfs(index+1, curr) {
			dp[index][curr] = 1
			return true
		}
		dp[index][curr] = 0
		return false
	}
	return dfs(0, 0)

}
