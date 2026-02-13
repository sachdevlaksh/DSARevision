/*
LeetCode Problem #121: Best Time to Buy and Sell Stock
Difficulty: Easy

You are given an array prices where prices[i] is the price of a given stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
*/

package LC121

import "math"


func maxProfit(prices []int) int {
	n := len(prices)

	maxProf := 0
	minPrice := math.MaxInt

	for i := 0; i < n; i++{
		minPrice = min(minPrice, prices[i])
		maxProf = max(maxProf, prices[i]-minPrice)
	}

	return maxProf
}