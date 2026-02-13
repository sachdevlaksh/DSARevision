/*
LeetCode Problem #5: Longest Palindromic Substring
Difficulty: Medium

Given a string s, return the longest palindromic substring in s.
*/

package LC5

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}
	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		l1, r1 := checkPalindrome(i, i, s)
		if r1-l1 > end-start {
			end = r1
			start = l1
		}
		l2, r2 := checkPalindrome(i, i+1, s)
		if r2-l2 > end-start {
			end = r2
			start = l2
		}
	}

	return s[start : end+1]
}

func checkPalindrome(left, right int, s string) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++

	}
	return left + 1, right - 1
}

