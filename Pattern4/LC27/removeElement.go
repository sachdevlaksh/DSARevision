/*
LeetCode Problem #27: Remove Element
Difficulty: Easy

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.
*/

package LC27

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
