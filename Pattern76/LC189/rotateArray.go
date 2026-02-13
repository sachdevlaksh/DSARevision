/*
LeetCode Problem #189: Rotate Array
Difficulty: Medium

Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/

package LC189

import "fmt"

func RotateArray(nums []int, k int) {
	n := len(nums)
	if k >= n {
		k = k % n
	}

	if len(nums) == 1 || k == 0 {
		return
	}
	nums2 := append(nums[n-k:], nums[:n-k]...)
	copy(nums, nums2)
	fmt.Println(nums)
}