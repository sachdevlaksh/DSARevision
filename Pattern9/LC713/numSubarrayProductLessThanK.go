/*
LeetCode Problem #713: Subarray Product Less Than K
Difficulty: Medium

Given an array of integers nums and an integer k, return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.
*/

package LC713

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	left := 0
	prod := 1
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for right >= left && prod >= k {
			prod = prod / nums[left]
			left++
		}
        count += right-left+1
	}
	return count
}
