/*
LeetCode Problem #16: 3Sum Closest
Difficulty: Medium

Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers.
*/

package LC16

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
sort.Ints(nums)
	closetSum := nums[0] + nums[1] + nums[2]
	for i, num := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
            sum := num + nums[l] + nums[r]
            if abs(sum-target) < abs(closetSum - target){
                closetSum  = sum
            }

			if sum < target {
				l++
			} else if sum > target {
				r--
			}else{
                return sum
            }
		}
	}
	return closetSum
}

func abs(num1 int) int {
	if num1 < 0 {
		return -1*num1
	} else {
		return num1
	}
}