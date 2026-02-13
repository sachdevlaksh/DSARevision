/*
LeetCode Problem #1438: Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
Difficulty: Medium

Given an array of integers nums and an integer limit, return the size of the longest contiguous subarray such that the absolute difference between any two elements is less than or equal to limit.
*/

package LC1438


func longestSubarray(nums []int, limit int) int {
	minLimit := -1
	for i := 0; i < len(nums);i++{
		if nums[i] <= limit{
			minLimit = 1
		}
		for j := i +1; j < len(nums); j++{
			if nums[j] - nums[i] < limit{
				return j-i
			}
		}
	}

	return minLimit
}