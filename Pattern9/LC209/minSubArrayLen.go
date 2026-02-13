/*
LeetCode Problem #209: Minimum Size Subarray Sum
Difficulty: Medium

Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target.
*/

package LC209

import "math"

func minSubArrayLen(target int, nums []int) int {
    mini := math.MaxInt
    sum := 0
	left := 0
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            mini = min(mini, right-left+1)
            sum -= nums[left]
            left++
        }
    }
    if mini == math.MaxInt {
        return 0
    }
    return mini
}
