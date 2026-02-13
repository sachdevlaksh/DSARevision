/*
LeetCode Problem #438: Find All Anagrams in a String
Difficulty: Medium

Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.
*/

package LC438

func findAnagrams(s string, p string) []int {

	var pmap [26]int
	for i := 0; i < len(p); i++ {
		pmap[p[i]-'a']++
	}
	res := make([]int, 0)
	missing := len(p)
	left := 0
	for right := 0; right < len(s); right++ {
		r := s[right] - 'a'
		if pmap[r] > 0 {
			missing--
		}
		pmap[s[right]]--
		if right-left+1 > len(p) {
			l := s[left] - 'a'
			if pmap[l] >= 0{
				missing++
			}
			pmap[s[left]]++
			left++
		}

		if missing == 0 {
			res = append(res, left)
		}
	}

	return res
}
