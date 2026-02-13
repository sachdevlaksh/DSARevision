/*
LeetCode Problem #3254: Find the Power of K-Size Subarrays - I
Difficulty: Medium

You are given an array of integers nums of size n, and a positive integer k. Find the power of each subarray of size k.
*/

package LC3254

func resultsArray(nums []int, k int) []int {

	if k <= 0 || k > len(nums) {
		return []int{}
	}

	out := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {

		flag := false
		for count := 1; count < k; count++ {
			if nums[i+count] != nums[i+count-1]+1 {
				out = append(out, -1)
				flag = true
				break
			}

		}
		if !flag {
			out = append(out, nums[i+k-1])
		}
	}

	return out
}
