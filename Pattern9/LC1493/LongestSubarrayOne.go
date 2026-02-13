/*
LeetCode Problem #1493: Longest Subarray of 1's After Removing One Element
Difficulty: Medium

Given a binary array nums, you should delete one element from it. Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if no such subarray exists.
*/

package LC1493

func LongestSubarray(nums []int) int {
	l, r, out, isZero := 0, 0, 0, false

	for _, num := range nums {
		if num == 1 {
			r++
		} else {
			l = r
			r = 0
			isZero = true
		}
		out = max(out, l+r)

	}
	if isZero {
		return out
	} else {
		return out - 1
	}

}
