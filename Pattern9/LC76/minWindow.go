/*
LeetCode Problem #76: Minimum Window Substring
Difficulty: Hard

Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t is included in the window.
*/

package LC76

import "math"

func minWindow(s string, t string) string {

	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	var need [256]int

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left := 0
	start := 0
	missing := len(t)
	minLength := math.MaxInt

	for right := 0; right < len(s); right++ {
		c := s[right]

		if need[c] > 0 {
			missing--
		}
		need[c]--

		for missing == 0 {
			windowLen := right - left + 1
			if windowLen < minLength {
				minLength = windowLen
				start = left
			}
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++
			}
			left++
		}
	}

	if minLength == math.MaxInt {
		return ""
	}
	return s[start : start+minLength]
}