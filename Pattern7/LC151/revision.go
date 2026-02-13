/*
LeetCode Problem #151: Reverse Words in a String
Difficulty: Medium

Given an input string s, reverse the order of the words. A word is defined as a sequence of non-space characters.
*/

package LC151

import "strings"

func reverseWords(s string) string {

	n := len(s)

	if n == 0 {
		return s
	}

	//To remove prefix space
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}

	//To remove Suffix space
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}

	if len(s) == 0 {
		return ""
	}
	end := len(s) - 1
	var sb strings.Builder
	for end >= 0 {
		right := end
		for right >= 0 && s[right] != ' ' {
			right--
		}
		sb.WriteString(s[right+1 : end+1])
		sb.WriteByte(' ')
		end = right - 1
		for end >= 0 && s[end] == ' ' {
			end--
		}
	}

	res := sb.String()

	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	//to remove the last space
	return res
}
