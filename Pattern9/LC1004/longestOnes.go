/*
LeetCode Problem #1004: Max Consecutive Ones III
Difficulty: Medium

Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/

package LC1004

func longestOnes(nums []int, k int) int {
	maxLeng := 0
	zeroLen := 0
	left := 0
	for right := 0; right < len(nums);right++{
		if nums[right] == 0{
			zeroLen++
		}
		for zeroLen > 0{
			if nums[right] == 0{
				zeroLen--
			}
			left++

		}
		maxLeng = max(maxLeng, right-left+1)
	}
	return maxLeng
}