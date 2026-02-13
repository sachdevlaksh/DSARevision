/*
LeetCode Problem #131: Palindrome Partitioning
Difficulty: Medium

Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
*/

package LC131

func partition(s string) [][]string {

	out := make([][]string, 0)
	n := len(s)

	var isPal func(start int, end int) bool

	isPal = func(start int, end int) bool {
		for start <= end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	var dfs func(index int, li []string)

	dfs = func(index int, li []string) {
		if index == n {
			clone := make([]string, len(li))
			copy(clone,li)
			out = append(out, clone)
			return
		}
		for i := index; i < n; i++ {
			if isPal(index, i) {
				li = append(li, s[index:i+1])
				dfs(i+1, li)
				li = li[:len(li)-1]
			}

		}
	}
	dfs(0, make([]string, 0))
	return out
}
