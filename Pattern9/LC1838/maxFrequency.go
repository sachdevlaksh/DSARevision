/*
LeetCode Problem #1838: Frequency of the Most Frequent Element
Difficulty: Medium

The frequency of an element is the number of times it appears in an array. You are given an integer array nums and an integer k. In one operation, you can choose an element of nums and increment it by 1.
*/

package LC1838

import (
	"slices"
)

func maxFrequency(nums []int, k int) int {
	maxLenth := 1
	slices.Sort(nums)
	sum := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for (right-left+1)*nums[right]-sum > k {
			sum -= nums[left]
			left++
		}
		maxLenth = max(maxLenth, right-left+1)
	}
	return maxLenth
}
