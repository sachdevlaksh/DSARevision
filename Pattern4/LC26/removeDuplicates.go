/*
LeetCode Problem #26: Remove Duplicates from Sorted Array
Difficulty: Easy

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. Return the number of unique elements in nums.
*/

package LC26

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
        return 0
    }
	count := 1
	for i := 1; i < len(nums); i++ {
		 if nums[i] != nums[i-1] {
            nums[count] = nums[i]
            count++
        }
	}
	return count
}
