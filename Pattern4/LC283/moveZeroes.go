/*
LeetCode Problem #283: Move Zeroes
Difficulty: Easy

Given an integer array nums, move all 0's to the end while maintaining the relative order of the non-zero elements. Note that you must do this in-place without making a copy of the array.
*/

package LC283

func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	for pos < len(nums){
		nums[pos] = 0
		pos++
	}
}
