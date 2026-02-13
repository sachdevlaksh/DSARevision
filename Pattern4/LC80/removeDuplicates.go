/*
LeetCode Problem #80: Remove Duplicates from Sorted Array II
Difficulty: Medium

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice.
*/

package LC80

func removeDuplicates(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}
	count := 2
	for i := 2; i < len(nums); i++ {
		 if nums[i] != nums[count-2] {
            nums[count] = nums[i]
            count++
        }
	}

	return count
}