/*
LeetCode Problem #152: Maximum Product Subarray
Difficulty: Medium

Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return that product.
*/

package LC152

import "math"

func MaxProduct(nums []int) int {
	n := len(nums)
	pre := 1
	suff := 1

	maxProd := math.MinInt64

	for i := 0; i < n; i++ {
		if pre == 0 {
			pre = 1
		}
		if suff == 0 {
			suff = 1
		}
		pre = pre * nums[i]
		suff = suff * nums[n-i-1]
		maxProd = max(maxProd, max(pre, suff))
	}
	return maxProd
}
