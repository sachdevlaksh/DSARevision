/*
LeetCode Problem #322: Coin Change
Difficulty: Medium

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money. Return the fewest number of coins that you need to make up that amount.
*/

package LC322

import "math"

func coinChange(coins []int, amount int) int {
	var rec func(amt int) int

	rec = func(amt int) int {
		if amt == 0 {
			return 0
		}
		if amt < 0 {
			return math.MaxInt
		}
		mini := math.MaxInt
		for _, coinCh := range coins {
			ans := rec(amt - coinCh)
			if ans != math.MaxInt {
				mini = min(mini, 1+ans)
			}
		}
		return mini

	}

	if ret := rec(amount); ret == math.MaxInt {
		return -1
	} else {
		return ret
	}

}

func coinChangeMem(coins []int, amount int) int {
	const INF = math.MaxInt / 2
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	var rec func(amt int, dp []int) int

	rec = func(amt int, dp []int) int {
		if amt == 0 {
			return 0
		}
		if amt < 0 {
			return INF
		}

		if dp[amt] != -1 {
			return dp[amt]
		}
		mini := INF
		for _, coinCh := range coins {
			ans := rec(amt-coinCh, dp)
			if ans != INF {
				mini = min(mini, 1+ans)
			}
		}
		dp[amt] = mini
		return mini

	}

	if ret := rec(amount, dp); ret == INF {
		return -1
	} else {
		return ret
	}

}
func coinChangeTab(coins []int, amount int) int {
	const INF = math.MaxInt / 2
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != INF {
				dp[i] = min(dp[i], 1+dp[i-coins[j]])
			}
		}
	}
	if dp[amount] == INF {
		return -1
	}
	return dp[amount]

}
