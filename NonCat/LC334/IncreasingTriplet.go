/*
LeetCode Problem #334: Increasing Triplet Subsequence
Difficulty: Medium

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
*/

package LC334

import (
	"fmt"
	"math"
)

func IncreasingTriplet(nums []int) bool {

	small, mid := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= small  {
			small = num
		} else if num < mid  {
			mid = num
		} else {
			fmt.Println(true)
			return true

		}
	}
	fmt.Println(false)
	return false
}
