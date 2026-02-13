/*
LeetCode Problem #214: Shortest Palindrome
Difficulty: Hard

You are given a string s. You can convert s to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can create by adding characters in front of s.
*/

package LC214

import (

	"strings"
)

func shortestPalindrome(s string) string {

	L := 0
	n := len(s)
	for i := n; i > 0; i-- {
		l, r := 0, i-1
		ok := true
		for l < r {
			if s[l] != s[r] {
				ok = false
				break
			}
			l++
			r--
		}
		if ok {
			L = i
			break
		}
	}

	if n == L {
		return s
	}
	var sb strings.Builder
	for i := n - 1; i >= L; i-- {
		sb.WriteByte(s[i])
	}
	sb.WriteString(s)
	return sb.String()
}
