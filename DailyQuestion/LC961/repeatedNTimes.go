/*
LeetCode Problem #961: N-Repeated Element in Size 2N Array
Difficulty: Easy

You are given an integer array nums with the following properties: nums.length == 2 * n. nums contains n + 1 unique elements. Exactly one element of nums is repeated n times. Return the element that is repeated n times.
*/

package LC961

func repeatedNTimes(nums []int) int {

	n := len(nums) / 2
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] == n {
			return num
		}
	}

	return -1
}
