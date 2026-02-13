/*
LeetCode Problem #424: Longest Repeating Character Replacement
Difficulty: Medium

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
*/

package LC424

func characterReplacement(s string, k int) int {
	maxF,left,res := 0,0,0
	li := make([]int, 26)
	for right := 0; right < len(s); right++{
		li[s[right]-'A']++
		maxF = max(maxF, li[s[right]-'A'])
		for right-left+1 > k + maxF{
			li[s[left]-'A']--
			left++
		}
		res = max(res,right-left+1)
	}
	return res
}