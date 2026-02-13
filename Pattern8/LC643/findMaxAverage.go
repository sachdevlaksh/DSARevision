/*
LeetCode Problem #643: Maximum Average Subarray I
Difficulty: Easy

You are given an integer array nums consisting of n elements, and an integer k. Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
*/

package LC643

import "math"

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 || k == 0 {
		return 0.0
	}
	i := 0
	maxAvg := math.Inf(-1)
	avg := float64(0)
	for i < len(nums)-k+1 {
		count := 0
		sum := 0
		for count < k {
			sum += nums[i+count]
			count++
		}
		avg = float64(sum) / float64(k)
		if avg > maxAvg {
			maxAvg = avg
		}
		i++
	}
	return maxAvg
}
