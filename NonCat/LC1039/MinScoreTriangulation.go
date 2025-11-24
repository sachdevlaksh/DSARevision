package LC1039

import "math"

func MinScoreTriangulation(values []int) int {

	n := len(values)
	//return score(values, 0, len(values)-1)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
	}

	return scoreMem(values, 0, n, dp)

}

func score(values []int, i, j int ) int {
	if i+1 == j {
		return 0
	}

	ans := math.MaxInt
	for k := i + 1; k < len(values); k++ {
		ans = min(ans, values[i]*values[j]*values[k]+score(values, i, k)+score(values, k, j))
	}
	return ans
}


func scoreMem(values []int, i, j int, dp [][]int ) int {
	if i+1 == j {
		return 0
	}

	if dp[i][j] != 0{
		return dp[i][j]
	}
	ans := math.MaxInt
	for k := i + 1; k < j; k++ {

		ans = min(ans, values[i]*values[j]*values[k]+scoreMem(values, i, k, dp)+scoreMem(values, k, j, dp))
	}
	dp[i][j] = ans
	return dp[i][j]
}