/*
LeetCode Problem #122: Best Time to Buy and Sell Stock II
Difficulty: Medium

You are given an integer array prices where prices[i] is the price of a given stock on the ith day. On each day, you may decide to buy and/or sell the stock.
*/

package LC122

func maxProfit(prices []int) int {
  profit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            profit += prices[i] - prices[i-1]
        }
    }
    return profit
}