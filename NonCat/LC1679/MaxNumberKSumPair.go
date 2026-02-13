/*
LeetCode Problem #1679: Max Number of K-Sum Pairs
Difficulty: Medium

You are given an integer array nums and an integer k. In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
*/

package main

import "sort"

func maxOperations(nums []int, k int) int {
	count := 0
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] < k {
			i++
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			count++
			i++
			j--
		}
	}
	return count
}
